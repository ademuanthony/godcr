!function(u){function e(e){for(var t,n,r=e[0],s=e[1],a=e[2],o=0,i=[];o<r.length;o++)n=r[o],d[n]&&i.push(d[n][0]),d[n]=0;for(t in s)Object.prototype.hasOwnProperty.call(s,t)&&(u[t]=s[t]);for(l&&l(e);i.length;)i.shift()();return h.push.apply(h,a||[]),c()}function c(){for(var e,t=0;t<h.length;t++){for(var n=h[t],r=!0,s=1;s<n.length;s++){var a=n[s];0!==d[a]&&(r=!1)}r&&(h.splice(t--,1),e=o(o.s=n[0]))}return e}var n={},d={0:0},h=[];function o(e){if(n[e])return n[e].exports;var t=n[e]={i:e,l:!1,exports:{}};return u[e].call(t.exports,t,t.exports,o),t.l=!0,t.exports}o.m=u,o.c=n,o.d=function(e,t,n){o.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:n})},o.r=function(e){"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},o.t=function(t,e){if(1&e&&(t=o(t)),8&e)return t;if(4&e&&"object"==typeof t&&t&&t.__esModule)return t;var n=Object.create(null);if(o.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var r in t)o.d(n,r,function(e){return t[e]}.bind(null,r));return n},o.n=function(e){var t=e&&e.__esModule?function(){return e.default}:function(){return e};return o.d(t,"a",t),t},o.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},o.p="../dist/";var t=window.webpackJsonp=window.webpackJsonp||[],r=t.push.bind(t);t.push=e,t=t.slice();for(var s=0;s<t.length;s++)e(t[s]);var l=r;h.push([10,1]),c()}({10:function(e,t,n){"use strict";n.r(t);var r=n(1),s=n(9),a=(n(11),r.a.start()),o=n(16);a.load(Object(s.a)(o))},11:function(e,t,n){var r=n(12);"string"==typeof r&&(r=[[e.i,r,""]]);var s={hmr:!0,transform:void 0,insertInto:void 0};n(14)(r,s);r.locals&&(e.exports=r.locals)},12:function(e,t,n){(e.exports=n(13)(!1)).push([e.i,"body {\n  background-color: #f5f5f5;\n  font-size: 13px; }\n\n.header-top .inner {\n  padding: 15px 0;\n  color: #333; }\n\n.navbar {\n  background-color: #fff;\n  box-shadow: 0 0 2px 2px rgba(0, 0, 0, 0.03), 1px 1px 2px 2px rgba(0, 0, 0, 0.03);\n  -webkit-box-shadow: 0 0 2px 2px rgba(0, 0, 0, 0.03), 1px 1px 2px 2px rgba(0, 0, 0, 0.03);\n  -moz-box-shadow: 0 0 2px 2px rgba(0, 0, 0, 0.03), 1px 1px 2px 2px rgba(0, 0, 0, 0.03); }\n\n.nav-link {\n  padding: 10px 22px !important; }\n\n.nav-link i,\n.nav-link span {\n  display: inline-block;\n  height: 100%;\n  font-size: 18px; }\n\n.content {\n  padding: 25px 0; }\n\n.table td,\n.table th {\n  padding: 0.3rem; }\n\n.table.inverse tbody td {\n  background-color: #f5f5f5 !important;\n  border-color: #fff !important; }\n\n.card {\n  border-radius: 0;\n  -moz-border-radius: 0;\n  -webkit-border-radius: 0; }\n\n.card-body {\n  padding: 1.3rem 1.8rem !important; }\n\n.form-control {\n  border-radius: 0; }\n\n.error {\n  color: brown; }\n\n.btn {\n  border-radius: 0;\n  -moz-border-radius: 0;\n  -webkit-border-radius: 0;\n  padding: 0.375rem 1rem; }\n\n.btn i {\n  font-size: 12px; }\n\n.modal-content {\n  border-radius: 0;\n  -moz-border-radius: 0;\n  -webkit-border-radius: 0; }\n\n.breadcrumb {\n  background: none;\n  padding: 1rem 0 0;\n  margin: 0; }\n\n.breadcrumb-item {\n  font-size: 15px; }\n\n.collapsible .card-header {\n  padding: 0 !important; }\n\n.card-header .btn-link {\n  text-decoration: none;\n  display: block;\n  cursor: pointer; }\n\n.no-btm-pad {\n  padding-bottom: 0 !important; }\n\nselect#source-account[disabled] {\n  border: none !important;\n  background-color: transparent !important;\n  padding: 0;\n  -webkit-appearance: none;\n  -moz-appearance: none; }\n\n/* -------------slide down animation------------- */\n.slide-down,\n.slide-up {\n  max-height: 0;\n  overflow-y: hidden;\n  transition: max-height 0.5s ease-in-out; }\n\n.slide-down {\n  max-height: 10em; }\n",""])},16:function(e,t,n){var r={"./purchase_tickets_controller.js":17,"./send_controller.js":37};function s(e){var t=a(e);return n(t)}function a(e){var t=r[e];if(t+1)return t;var n=new Error("Cannot find module '"+e+"'");throw n.code="MODULE_NOT_FOUND",n}s.keys=function(){return Object.keys(r)},s.resolve=a,(e.exports=s).id=16},17:function(e,t,n){"use strict";n.r(t),n.d(t,"default",function(){return h});var a=n(1),r=n(2),o=n.n(r);function s(e){return(s="function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(e){return typeof e}:function(e){return e&&"function"==typeof Symbol&&e.constructor===Symbol&&e!==Symbol.prototype?"symbol":typeof e})(e)}function i(e,t){for(var n=0;n<t.length;n++){var r=t[n];r.enumerable=r.enumerable||!1,r.configurable=!0,"value"in r&&(r.writable=!0),Object.defineProperty(e,r.key,r)}}function u(e,t){return!t||"object"!==s(t)&&"function"!=typeof t?function(e){if(void 0!==e)return e;throw new ReferenceError("this hasn't been initialised - super() hasn't been called")}(e):t}function c(e){return(c=Object.setPrototypeOf?Object.getPrototypeOf:function(e){return e.__proto__||Object.getPrototypeOf(e)})(e)}function d(e,t){return(d=Object.setPrototypeOf||function(e,t){return e.__proto__=t,e})(e,t)}var h=function(e){function t(){return function(e,t){if(!(e instanceof t))throw new TypeError("Cannot call a class as a function")}(this,t),u(this,c(t).apply(this,arguments))}var n,r,s;return function(e,t){if("function"!=typeof t&&null!==t)throw new TypeError("Super expression must either be null or a function");e.prototype=Object.create(t&&t.prototype,{constructor:{value:e,writable:!0,configurable:!0}}),t&&d(e,t)}(t,a["b"]),n=t,s=[{key:"targets",get:function(){return["sourceAccount","errors","spendUnconfirmed","numberOfTickets","walletPassphrase","successMessage","errorMessage","passwordError","submitButton"]}}],(r=[{key:"validateForm",value:function(){var e=[],t=!(this.errorsTarget.innerHTML="");if(""===this.sourceAccountTarget.value&&e.push("The source account is required"),""===this.numberOfTicketsTarget.value&&e.push("The number of tickets is required"),e.length){for(var n in e)this.showError(e[n]);t=!1}return t}},{key:"submitForm",value:function(){if(this.validatePassphrase()){$("#passphrase-modal").modal("hide"),this.submitButtonTarget.innerHTML="Purchasing...",this.submitButtonTarget.setAttribute("disabled","disabled");var e=$("#purchase-tickets-form").serialize();this.walletPassphraseTarget.value="",this.sourceAccountTarget.disabled&&(e+="&source-account="+this.sourceAccountTarget.value);var s=this;o.a.post("/purchase-tickets",e).then(function(e){var t=e.data;if(t.success){var n=["<p>You have purchased "+t.message.length+" ticket(s)</p>"],r=t.message.map(function(e){return"<p><strong>"+e+"</strong></p>"});n.push(r),s.setSuccessMessage(n.join(""))}else s.setErrorMessage(t.message)}).catch(function(){s.setErrorMessage("A server error occurred")}).then(function(){s.submitButtonTarget.innerHTML="Purchase",s.submitButtonTarget.removeAttribute("disabled")})}}},{key:"validatePassphrase",value:function(){return""!==this.walletPassphraseTarget.value||!(this.passwordErrorTarget.innerHTML='<div class="error">Your wallet passphrase is required</div>')}},{key:"getWalletPassphraseAndSubmit",value:function(){this.clearMessages(),this.validateForm()&&$("#passphrase-modal").modal()}},{key:"setErrorMessage",value:function(e){this.hide(this.successMessageTarget),this.show(this.errorMessageTarget),this.errorMessageTarget.innerHTML=e}},{key:"setSuccessMessage",value:function(e){this.hide(this.errorMessageTarget),this.show(this.successMessageTarget),this.successMessageTarget.innerHTML=e}},{key:"clearMessages",value:function(){this.hide(this.errorMessageTarget),this.hide(this.successMessageTarget),this.errorsTarget.innerHTML=""}},{key:"hide",value:function(e){e.classList.add("d-none")}},{key:"show",value:function(e){e.classList.remove("d-none")}},{key:"showError",value:function(e){this.errorsTarget.innerHTML+='<div class="error">'.concat(e,"</div>")}}])&&i(n.prototype,r),s&&i(n,s),t}()},37:function(e,t,n){"use strict";n.r(t),n.d(t,"default",function(){return h});var a=n(1),r=n(2),o=n.n(r);function s(e){return(s="function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(e){return typeof e}:function(e){return e&&"function"==typeof Symbol&&e.constructor===Symbol&&e!==Symbol.prototype?"symbol":typeof e})(e)}function i(e,t){for(var n=0;n<t.length;n++){var r=t[n];r.enumerable=r.enumerable||!1,r.configurable=!0,"value"in r&&(r.writable=!0),Object.defineProperty(e,r.key,r)}}function u(e,t){return!t||"object"!==s(t)&&"function"!=typeof t?function(e){if(void 0!==e)return e;throw new ReferenceError("this hasn't been initialised - super() hasn't been called")}(e):t}function c(e){return(c=Object.setPrototypeOf?Object.getPrototypeOf:function(e){return e.__proto__||Object.getPrototypeOf(e)})(e)}function d(e,t){return(d=Object.setPrototypeOf||function(e,t){return e.__proto__=t,e})(e,t)}var h=function(e){function t(){return function(e,t){if(!(e instanceof t))throw new TypeError("Cannot call a class as a function")}(this,t),u(this,c(t).apply(this,arguments))}var n,r,s;return function(e,t){if("function"!=typeof t&&null!==t)throw new TypeError("Super expression must either be null or a function");e.prototype=Object.create(t&&t.prototype,{constructor:{value:e,writable:!0,configurable:!0}}),t&&d(e,t)}(t,a["b"]),n=t,s=[{key:"targets",get:function(){return["sourceAccount","address","amount","destinations","destinationTemplate","changeAddress","changeAmount","errors","customInput","customTxRow","customInputContent","submitButton","nextButton","removeDestinationButton","form","walletPassphrase","passwordError","useCustom","spendUnconfirmed","errorMessage","successMessage","progressBar","changeOutputsCard","changeOutputPnl","numberOfChangeOutputs","useRandomChangeOutputs","generateOutputsButton","changeOutputContent","changeDestinationTemplate","changeOutputAddress","changeOutputAmount"]}}],(r=[{key:"initialize",value:function(){this.newDestination()}},{key:"validateSendForm",value:function(e){this.errorsTarget.innerHTML="";var t=[],n=this.validateDestinationsField()&&(e||this.validateChangeOutputField());if(""===this.sourceAccountTarget.value&&t.push("The source account is required"),this.useCustomTarget.value&&this.getSelectedInputsSum()<this.getTotalSendAmount()&&t.push("The sum of selected inputs is less than send amount"),t.length){for(var r in t)this.showError(t[r]);n=!1}return n}},{key:"validateAndRefreshPercentage",value:function(){this.validateDestinationsField()&&this.calculateSelectedInputPercentage()}},{key:"validateDestinationsField",value:function(){var n=this;this.clearMessages();var r=!0;return this.addressTargets.forEach(function(e,t){""===e.value&&(n.showError("The destination address and amount are required"),r=!1)}),this.amountTargets.forEach(function(e,t){0<parseFloat(e.value)||(n.showError("Amount must be a non-zero positive number"),r=!1)}),r}},{key:"validateChangeOutputField",value:function(){var n=this;this.clearMessages();var r=!0;return this.changeOutputAmountTargets.forEach(function(e,t){0<parseFloat(e.value)||(n.showError("Change amount must be a non-zero positive number"),r=!1)}),r}},{key:"getSelectedInputsSum",value:function(){var n=0;return this.customTxRowTarget.querySelectorAll("input.custom-input:checked").forEach(function(e,t){n+=parseFloat(e.dataset.amount)}),n}},{key:"getTotalSendAmount",value:function(){var n=0;return this.amountTargets.forEach(function(e,t){n+=parseFloat(e.value)}),n}},{key:"calculateSelectedInputPercentage",value:function(){var e=this.getTotalSendAmount(),t=this.getSelectedInputsSum(),n=0;n=e<=t?100:t/e*100,this.progressBarTarget.style.width="".concat(n,"%")}},{key:"openCustomizePanel",value:function(){var n=this;this.resetCustomizePanel(),$("#custom-tx-row").slideDown();var e=n.sourceAccountTarget.value;this.getUnspentOutputs(e,function(e){var t=e.map(function(e){var t=new Date(1e3*e.receive_time).toString().split(" ").slice(0,5).join(" "),n=e.amount/1e8;return"<tr>\n                  <td width='5%'>\n                    <input data-action='click->send#calculateSelectedInputPercentage' type='checkbox' class='custom-input' name='utxo' value='".concat(e.key,"' data-amount='").concat(n,"' />\n                  </td>\n                  <td width='40%'>").concat(e.address,"</td>\n                  <td width='15%'>").concat(n," DCR</td>\n                  <td width='25%'>").concat(t,"</td>\n                  <td width='15%'>").concat(e.confirmations," confirmation(s)</td>\n                </tr>")}).join("");n.customInputContentTarget.innerHTML=t,$("#custom-tx-row .status").hide()})}},{key:"getUnspentOutputs",value:function(e,n){this.nextButtonTarget.innerHTML="Loading...",this.nextButtonTarget.setAttribute("disabled","disabled");var r=this,t="/unspent-outputs/".concat(e);r.spendUnconfirmedTarget.checked&&(t+="?getUnconfirmed=true"),o.a.get(t).then(function(e){var t=e.data;t.success?n(t.message):r.setErrorMessage(t.message)}).catch(function(){r.setErrorMessage("A server error occurred")}).then(function(){r.nextButtonTarget.innerHTML="Next",r.nextButtonTarget.removeAttribute("disabled")})}},{key:"toggleUseRandomChangeOutputs",value:function(){0<parseFloat(this.numberOfChangeOutputsTarget.value)&&this.generateChangeOutputs()}},{key:"changeOutputNumberChanged",value:function(){0<parseFloat(this.numberOfChangeOutputsTarget.value)&&this.generateChangeOutputs()}},{key:"generateChangeOutputs",value:function(){if(this.validateSendForm(!0)&&!this.generatingChangeOutputs){this.generatingChangeOutputs=!0,this.changeOutputContentTarget.innerHTML="";var e=parseFloat(this.numberOfChangeOutputsTarget.value);if(0<e){this.generateOutputsButtonTarget.setAttribute("disabled","disabled"),this.generateOutputsButtonTarget.innerHTML="Loading...",this.numberOfChangeOutputsTarget.setAttribute("disabled","disabled");var n=this;n.getRandomChangeOutputs(e,function(e){e.forEach(function(e){var t=document.importNode(n.changeDestinationTemplateTarget.content,!0);t.querySelector('input[name="change-output-address"]').value=e.Address,n.useRandomChangeOutputsTarget.checked&&(t.querySelector('input[name="change-output-amount"]').value=e.Amount),n.changeOutputContentTarget.appendChild(t)}),n.show(n.changeOutputContentTarget),n.useRandomChangeOutputsTarget.checked&&n.changeOutputAmountTargets.forEach(function(e){e.setAttribute("readonly","true")}),n.generateOutputsButtonTarget.removeAttribute("disabled"),n.generateOutputsButtonTarget.innerHTML="Generate Change Outputs",n.numberOfChangeOutputsTarget.removeAttribute("disabled"),n.generatingChangeOutputs=!1},function(){n.generateOutputsButtonTarget.removeAttribute("disabled"),n.generateOutputsButtonTarget.innerHTML="Generate Change Outputs",n.numberOfChangeOutputsTarget.removeAttribute("disabled"),n.generatingChangeOutputs=!1})}else this.showError("Number of change outputs must be a non-zero positive number")}}},{key:"getRandomChangeOutputs",value:function(e,n,t){var r=$("#send-form").serialize();r+="&totalSelectedInputAmountDcr=".concat(this.getSelectedInputsSum()),this.sourceAccountTarget.disabled&&(r+="&source-account=".concat(this.sourceAccountTarget.value)),r+="&nChangeOutput=".concat(e);var s=this;o.a.get("/random-change-outputs?"+r).then(function(e){var t=e.data;void 0!==t.error?s.setErrorMessage(t.error):n(t.message)}).catch(function(e){console.log(e),s.setErrorMessage("A server error occurred")}).then(function(){t&&t()})}},{key:"submitForm",value:function(){if(this.validatePassphrase()){$("#passphrase-modal").modal("hide"),this.nextButtonTarget.innerHTML="Sending...",this.nextButtonTarget.setAttribute("disabled","disabled");var e=$("#send-form").serialize();e+="&totalSelectedInputAmountDcr=".concat(this.getSelectedInputsSum()),this.walletPassphraseTarget.value="",this.sourceAccountTarget.disabled&&(e+="&source-account=".concat(this.sourceAccountTarget.value));var r=this;o.a.post("/send",e).then(function(e){var t=e.data;if(void 0!==t.error)r.setErrorMessage(t.error);else{r.resetSendForm();var n="The transaction was published successfully. Hash: <strong>".concat(t.txHash,"</strong>");r.setSuccessMessage(n)}}).catch(function(){r.setErrorMessage("A server error occurred")}).then(function(){r.nextButtonTarget.innerHTML="Next",r.nextButtonTarget.removeAttribute("disabled")})}}},{key:"resetSendForm",value:function(){for(this.resetCustomizePanel();1<this.destinationCount();)this.removeDestination();this.addressTargets.forEach(function(e){e.value=""}),this.amountTargets.forEach(function(e){e.value=""}),this.spendUnconfirmedTarget.checked=!1,this.useCustomTarget.checked=!1,$("#custom-tx-row").slideUp(),this.hide(this.changeOutputsCardTarget),this.clearMessages()}},{key:"resetCustomizePanel",value:function(){var t=this;this.customTxRowTarget.querySelectorAll("tbody").forEach(function(e){e.innerHTML=""}),this.customTxRowTarget.querySelectorAll(".status").forEach(function(e){t.show(e)}),this.customTxRowTarget.querySelectorAll(".alert-danger").forEach(function(e){e.parentNode.removeChild(e)}),this.hide(this.changeOutputsCardTarget),this.useRandomChangeOutputsTarget.checked=!1,this.numberOfChangeOutputsTarget.value="",this.changeOutputContentTarget.innerHTML=""}},{key:"newDestination",value:function(){if(this.validateDestinationsField()){var e=document.importNode(this.destinationTemplateTarget.content,!0);this.destinationsTarget.appendChild(e),1<this.destinationCount()&&this.show(this.removeDestinationButtonTarget)}}},{key:"removeDestination",value:function(){1<this.destinationsTarget.querySelectorAll("div.destination").length&&(this.destinationsTarget.removeChild(this.destinationsTarget.querySelector("div.destination:last-child")),1<this.destinationCount()||this.hide(this.removeDestinationButtonTarget))}},{key:"destinationCount",value:function(){return this.destinationsTarget.querySelectorAll("div.destination").length}},{key:"validatePassphrase",value:function(){return""!==this.walletPassphraseTarget.value||!(this.passwordErrorTarget.innerHTML='<div class="error">Your wallet passphrase is required</div>')}},{key:"getWalletPassphraseAndSubmit",value:function(){this.clearMessages(),this.validateDestinationsField()&&$("#passphrase-modal").modal()}},{key:"toggleUseCustom",value:function(){if(!this.useCustomTarget.checked)return $("#custom-tx-row").slideUp(),this.resetCustomizePanel(),void this.hide(this.changeOutputsCardTarget);this.validateDestinationsField()?(this.clearMessages(),this.openCustomizePanel(),this.show(this.changeOutputsCardTarget)):this.useCustomTarget.checked=!1}},{key:"toggleSpendUnconfirmed",value:function(){this.useCustomTarget.checked&&(this.resetCustomizePanel(),this.openCustomizePanel())}},{key:"setErrorMessage",value:function(e){this.hide(this.successMessageTarget),this.show(this.errorMessageTarget),this.errorMessageTarget.innerHTML=e}},{key:"setSuccessMessage",value:function(e){this.hide(this.errorMessageTarget),this.show(this.successMessageTarget),this.successMessageTarget.innerHTML=e}},{key:"clearMessages",value:function(){this.hide(this.errorMessageTarget),this.hide(this.successMessageTarget),this.errorsTarget.innerHTML="",this.successMessageTarget.innerHTML=""}},{key:"hide",value:function(e){e.classList.add("d-none")}},{key:"show",value:function(e){e.classList.remove("d-none")}},{key:"showError",value:function(e){this.errorsTarget.innerHTML+='<div class="error">'.concat(e,"</div>")}}])&&i(n.prototype,r),s&&i(n,s),t}()}});
//# sourceMappingURL=app.bundle.js.map