(window.webpackJsonp=window.webpackJsonp||[]).push([[2],{50:function(n,t,o){(n.exports=o(14)(!1)).push([n.i,'/*\\\n|*| ========================================================================\n|*| Bootstrap Toggle: bootstrap4-toggle.css v3.4.0\n|*| https://gitbrent.github.io/bootstrap-toggle/\n|*| ========================================================================\n|*| Copyright 2018-2019 Brent Ely\n|*| Licensed under MIT\n|*| ========================================================================\n\\*/\n\n/*\n* @added 3.0.0: Return support for "*-xs" removed in Bootstrap-4\n* @see: [Comment](https://github.com/twbs/bootstrap/issues/21881#issuecomment-341972830)\n*/\n.btn-group-xs > .btn, .btn-xs {\n\tpadding: .35rem .4rem .25rem .4rem;\n\tfont-size: .875rem;\n\tline-height: .5;\n\tborder-radius: .2rem;\n}\n\n.checkbox label .toggle, .checkbox-inline .toggle {\n\tmargin-left: -1.25rem;\n\tmargin-right: .35rem;\n}\n\n.toggle {\n\tposition: relative;\n\toverflow: hidden;\n}\n.toggle.btn.btn-light, .toggle.btn.btn-outline-light {\n\t/* bootstrap-4 - add a border so toggle is delineated */\n\tborder-color: rgba(0, 0, 0, .15);\n}\n.toggle input[type="checkbox"] {\n\tdisplay: none;\n}\n.toggle-group {\n\tposition: absolute;\n\twidth: 200%;\n\ttop: 0;\n\tbottom: 0;\n\tleft: 0;\n\ttransition: left 0.35s;\n\t-webkit-transition: left 0.35s;\n\t-moz-user-select: none;\n\t-webkit-user-select: none;\n}\n.toggle.off .toggle-group {\n\tleft: -100%;\n}\n.toggle-on {\n\tposition: absolute;\n\ttop: 0;\n\tbottom: 0;\n\tleft: 0;\n\tright: 50%;\n\tmargin: 0;\n\tborder: 0;\n\tborder-radius: 0;\n}\n.toggle-off {\n\tposition: absolute;\n\ttop: 0;\n\tbottom: 0;\n\tleft: 50%;\n\tright: 0;\n\tmargin: 0;\n\tborder: 0;\n\tborder-radius: 0;\n\tbox-shadow: none; /* Bootstrap 4.0 Support via (Issue #186)[https://github.com/minhur/bootstrap-toggle/issues/186]) */\n}\n.toggle-handle {\n\tposition: relative;\n\tmargin: 0 auto;\n\tpadding-top: 0px;\n\tpadding-bottom: 0px;\n\theight: 100%;\n\twidth: 0px;\n\tborder-width: 0 1px;\n\tbackground-color: #fff;\n}\n\n.toggle.btn-outline-primary .toggle-handle {\n\tbackground-color: var(--primary);\n\tborder-color: var(--primary);\n}\n.toggle.btn-outline-secondary .toggle-handle {\n\tbackground-color: var(--secondary);\n\tborder-color: var(--secondary);\n}\n.toggle.btn-outline-success .toggle-handle {\n\tbackground-color: var(--success);\n\tborder-color: var(--success);\n}\n.toggle.btn-outline-danger .toggle-handle {\n\tbackground-color: var(--danger);\n\tborder-color: var(--danger);\n}\n.toggle.btn-outline-warning .toggle-handle {\n\tbackground-color: var(--warning);\n\tborder-color: var(--warning);\n}\n.toggle.btn-outline-info .toggle-handle {\n\tbackground-color: var(--info);\n\tborder-color: var(--info);\n}\n.toggle.btn-outline-light .toggle-handle {\n\tbackground-color: var(--light);\n\tborder-color: var(--light);\n}\n.toggle.btn-outline-dark .toggle-handle {\n\tbackground-color: var(--dark);\n\tborder-color: var(--dark);\n}\n.toggle[class*="btn-outline"]:hover .toggle-handle {\n\tbackground-color: var(--light);\n\topacity: 0.5;\n}\n\n/* NOTE: Must come first, so classes below override as needed */\n/* [default] (bootstrap-4.1.3 - .btn - h:38px) */\n.toggle.btn { min-width: 3.7rem; min-height: 2.15rem; }\n.toggle-on.btn { padding-right: 1.5rem; }\n.toggle-off.btn { padding-left: 1.5rem; }\n\n/* `lg` (bootstrap-4.1.3 - .btn - h:48px) */\n.toggle.btn-lg { min-width: 5rem; min-height: 2.815rem; }\n.toggle-on.btn-lg { padding-right: 2rem; }\n.toggle-off.btn-lg { padding-left: 2rem; }\n.toggle-handle.btn-lg { width: 2.5rem; }\n\n/* `sm` (bootstrap-4.1.3 - .btn - h:31px) */\n.toggle.btn-sm { min-width: 3.125rem; min-height: 1.938rem; }\n.toggle-on.btn-sm { padding-right: 1rem; }\n.toggle-off.btn-sm { padding-left: 1rem; }\n\n/* `xs` (bootstrap-3.3 - .btn - h:22px) */\n.toggle.btn-xs { min-width: 2.19rem; min-height: 1.375rem; }\n.toggle-on.btn-xs { padding-right: .8rem; }\n.toggle-off.btn-xs { padding-left: .8rem; }\n',""])},54:function(n,t,o){var e=o(50);"string"==typeof e&&(e=[[n.i,e,""]]);var r={hmr:!0,transform:void 0,insertInto:void 0};o(15)(e,r);e.locals&&(n.exports=e.locals)}}]);
//# sourceMappingURL=2.bundle.js.map