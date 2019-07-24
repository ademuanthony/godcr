package fyne

import (
	"fmt"
	"github.com/raedahgroup/godcr/fyne/pages"
	"strconv"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/raedahgroup/dcrlibwallet/defaultsynclistener"
	godcrApp "github.com/raedahgroup/godcr/app"
	"github.com/raedahgroup/godcr/fyne/widgets"
)

func ShowSyncWindow(wallet godcrApp.WalletMiddleware, window fyne.Window, App fyne.App) fyne.CanvasObject {
	progressBar := widget.NewProgressBar()
	progressBar.Min = 0
	progressBar.Max = 100

	var fullSyncReport string

	reportLabel := widget.NewLabel("")
	widget.Refresh(reportLabel)
	reportLabel.Hide()
	reportLabel.Alignment = fyne.TextAlignCenter
	var infoButton *widget.Button

	infoButton = widget.NewButton("Tap to view informations", func() {
		if infoButton.Text == "Tap to view informations" {

			widget.Refresh(reportLabel)
			reportLabel.Show()
			infoButton.SetText("Tap to hide informations")

		} else {
			widget.Refresh(reportLabel)
			reportLabel.Hide()
			infoButton.SetText("Tap to view informations")
		}
	})

	var syncDone bool

	wallet.SyncBlockChain(false, func(report *defaultsynclistener.ProgressReport) {
		progressReport := report.Read()
		progressBar.SetValue(float64(progressReport.TotalSyncProgress))

		if progressReport.Status == defaultsynclistener.SyncStatusSuccess {
			if syncDone == false {
				window.SetContent(pages.Menu(wallet))
			}
			syncDone = true
		}

		stringReport := strings.Builder{}
		if progressReport.TotalTimeRemaining == "" {
			stringReport.WriteString(fmt.Sprintf("%d%% completed.\n", progressReport.TotalSyncProgress))
		} else {
			stringReport.WriteString(fmt.Sprintf("%d%% completed, %s remaining.\n", progressReport.TotalSyncProgress, progressReport.TotalTimeRemaining))
		}

		reportLabel.SetText(strings.TrimSpace(stringReport.String()))
		widget.Refresh(reportLabel)

		switch progressReport.CurrentStep {
		case defaultsynclistener.FetchingBlockHeaders:
			stringReport.WriteString(fmt.Sprintf("Fetched %d of %d block headers.\n", progressReport.FetchedHeadersCount, progressReport.TotalHeadersToFetch))
			widget.Refresh(pages.BlkHeight)
			pages.BlkHeight.SetText(strconv.Itoa(int(progressReport.CurrentRescanHeight)) + " Blocks Connected")
			widget.Refresh(pages.BlkHeight)
			stringReport.WriteString(fmt.Sprintf("%d%% through step 1 of 3.\n", progressReport.HeadersFetchProgress))

			if progressReport.DaysBehind != "" {
				stringReport.WriteString(fmt.Sprintf("Your wallet is %s behind.\n", progressReport.DaysBehind))
			}
		case defaultsynclistener.DiscoveringUsedAddresses:
			stringReport.WriteString("Discovering used addresses.\n")
			if progressReport.AddressDiscoveryProgress > 100 {
				stringReport.WriteString(fmt.Sprintf("%d%% (over) through step 2 of 3.\n", progressReport.AddressDiscoveryProgress))
			} else {
				stringReport.WriteString(fmt.Sprintf("%d%% through step 2 of 3.\n", progressReport.AddressDiscoveryProgress))
			}

		case defaultsynclistener.ScanningBlockHeaders:
			stringReport.WriteString(fmt.Sprintf("Scanning %d of %d block headers.\n", progressReport.CurrentRescanHeight,
				progressReport.TotalHeadersToFetch))
			widget.Refresh(pages.BlkHeight)
			pages.BlkHeight.SetText(strconv.Itoa(int(progressReport.CurrentRescanHeight)) + " Blocks Connected")
			widget.Refresh(pages.BlkHeight)
			stringReport.WriteString(fmt.Sprintf("%d%% through step 3 of 3.\n", progressReport.HeadersFetchProgress))
		}

		// show peer count last
		netType := wallet.NetType()
		if progressReport.ConnectedPeers == 1 {
			stringReport.WriteString(fmt.Sprintf("Syncing with %d peer on %s.\n", progressReport.ConnectedPeers, netType))
		} else {
			stringReport.WriteString(fmt.Sprintf("Syncing with %d peers on %s.\n", progressReport.ConnectedPeers, netType))
		}
		if progressReport.ConnectedPeers <= 1 {
			widget.Refresh(pages.PeerConn)
			pages.PeerConn.SetText(strconv.Itoa(int(progressReport.ConnectedPeers)) + " Peer Connected")
			widget.Refresh(pages.PeerConn)
		} else {
			widget.Refresh(pages.PeerConn)
			pages.PeerConn.SetText(strconv.Itoa(int(progressReport.ConnectedPeers)) + " Peers Connected")
			widget.Refresh(pages.PeerConn)
		}

		fullSyncReport = stringReport.String()
		reportLabel.SetText(fullSyncReport)
		widget.Refresh(reportLabel)
	})

	return widget.NewVBox(
		widgets.NewVSpacer(10),
		widget.NewLabelWithStyle("Synchronizing....", fyne.TextAlignLeading, fyne.TextStyle{Italic: true, Bold: true}),
		widgets.NewVSpacer(10),
		progressBar,
		fyne.NewContainerWithLayout(layout.NewFixedGridLayout(infoButton.MinSize()), infoButton),
		reportLabel)
}
