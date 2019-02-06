package ginception

const (
	exceptionPage = `<!DOCTYPE html>
<!-- saved from url=(0023)https://localhost:5001/ -->
<html lang="pl" xmlns="http://www.w3.org/1999/xhtml"><head><meta http-equiv="Content-Type" content="text/html; charset=UTF-8"><style class="darkreader darkreader--fallback" media="screen"></style><style class="darkreader darkreader--text" media="screen"></style><style class="darkreader darkreader--invert" media="screen"></style><style class="darkreader darkreader--inline" media="screen">[data-darkreader-inline-bgcolor] {
  background-color: var(--darkreader-inline-bgcolor) !important;
}
[data-darkreader-inline-bgimage] {
  background-image: var(--darkreader-inline-bgimage) !important;
}
[data-darkreader-inline-border] {
  border-color: var(--darkreader-inline-border) !important;
}
[data-darkreader-inline-border-bottom] {
  border-bottom-color: var(--darkreader-inline-border-bottom) !important;
}
[data-darkreader-inline-border-left] {
  border-left-color: var(--darkreader-inline-border-left) !important;
}
[data-darkreader-inline-border-right] {
  border-right-color: var(--darkreader-inline-border-right) !important;
}
[data-darkreader-inline-border-top] {
  border-top-color: var(--darkreader-inline-border-top) !important;
}
[data-darkreader-inline-boxshadow] {
  box-shadow: var(--darkreader-inline-boxshadow) !important;
}
[data-darkreader-inline-color] {
  color: var(--darkreader-inline-color) !important;
}
[data-darkreader-inline-fill] {
  fill: var(--darkreader-inline-fill) !important;
}
[data-darkreader-inline-stroke] {
  stroke: var(--darkreader-inline-stroke) !important;
}
[data-darkreader-inline-outline] {
  outline-color: var(--darkreader-inline-outline) !important;
}</style><style class="darkreader darkreader--user-agent" media="screen">html {
    background-color: #17181c !important;
}
html, body, input, textarea, select, button {
    background-color: #17181c;
}
html, body, input, textarea, select, button {
    border-color: #575757;
    color: #e8e7e3;
}
a {
    color: #3391ff;
}
table {
    border-color: #4c4c4c;
}
::placeholder {
    color: #b9b4ac;
}
::selection {
    background-color: #005ccc;
    color: #ffffff;
}
::-moz-selection {
    background-color: #005ccc;
    color: #ffffff;
}
input:-webkit-autofill,
textarea:-webkit-autofill,
select:-webkit-autofill {
    background-color: #545b00 !important;
    color: #e8e7e3 !important;
}
::-webkit-scrollbar {
    background-color: #1b1c21;
    color: #c4c0ba;
}
::-webkit-scrollbar-thumb {
    background-color: #252a33;
}
::-webkit-scrollbar-thumb:hover {
    background-color: #2a323e;
}
::-webkit-scrollbar-thumb:active {
    background-color: #303d4f;
}
::-webkit-scrollbar-corner {
    background-color: #17181c;
}
* {
    scrollbar-color: #252a33 #1b1c21;
}</style>
        
        <title>Internal Server Error</title>
        <style>
            body {
    font-family: 'Segoe UI', Tahoma, Arial, Helvetica, sans-serif;
    font-size: .813em;
    color: #222;
    background-color: #fff;
}

h1, h2, h3, h4, h5 {
    /*font-family: 'Segoe UI',Tahoma,Arial,Helvetica,sans-serif;*/
    font-weight: 100;
}

h1 {
    color: #44525e;
    margin: 15px 0 15px 0;
}

h2 {
    margin: 10px 5px 0 0;
}

h3 {
    color: #363636;
    margin: 5px 5px 0 0;
}

code {
    font-family: Consolas, "Courier New", courier, monospace;
}

body .titleerror {
    padding: 3px 3px 6px 3px;
    display: block;
    font-size: 1.5em;
    font-weight: 100;
}

body .location {
    margin: 3px 0 10px 30px;
}

#header {
    font-size: 18px;
    padding: 15px 0;
    border-top: 1px #ddd solid;
    border-bottom: 1px #ddd solid;
    margin-bottom: 0;
}

    #header li {
        display: inline;
        margin: 5px;
        padding: 5px;
        color: #a0a0a0;
        cursor: pointer;
    }

    #header .selected {
        background: #44c5f2;
        color: #fff;
    }

#stackpage ul {
    list-style: none;
    padding-left: 0;
    margin: 0;
    /*border-bottom: 1px #ddd solid;*/
}

#stackpage .details {
    font-size: 1.2em;
    padding: 3px;
    color: #000;
}

#stackpage .stackerror {
    padding: 5px;
    border-bottom: 1px #ddd solid;
}


#stackpage .frame {
    padding: 0;
    margin: 0 0 0 30px;
}

    #stackpage .frame h3 {
        padding: 2px;
        margin: 0;
    }

#stackpage .source {
    padding: 0 0 0 30px;
}

    #stackpage .source ol li {
        font-family: Consolas, "Courier New", courier, monospace;
        white-space: pre;
        background-color: #fbfbfb;
    }

#stackpage .frame .source .highlight li span {
    color: #FF0000;
}

#stackpage .source ol.collapsible li {
    color: #888;
}

    #stackpage .source ol.collapsible li span {
        color: #606060;
    }

.page table {
    border-collapse: separate;
    border-spacing: 0;
    margin: 0 0 20px;
}

.page th {
    vertical-align: bottom;
    padding: 10px 5px 5px 5px;
    font-weight: 400;
    color: #a0a0a0;
    text-align: left;
}

.page td {
    padding: 3px 10px;
}

.page th, .page td {
    border-right: 1px #ddd solid;
    border-bottom: 1px #ddd solid;
    border-left: 1px transparent solid;
    border-top: 1px transparent solid;
    box-sizing: border-box;
}

    .page th:last-child, .page td:last-child {
        border-right: 1px transparent solid;
    }

.page .length {
    text-align: right;
}

a {
    color: #1ba1e2;
    text-decoration: none;
}

    a:hover {
        color: #13709e;
        text-decoration: underline;
    }

.showRawException {
    cursor: pointer;
    color: #44c5f2;
    background-color: transparent;
    font-size: 1.2em;
    text-align: left;
    text-decoration: none;
    display: inline-block;
    border: 0;
    padding: 0;
}

.rawExceptionStackTrace {
    font-size: 1.2em;
}

.rawExceptionBlock {
    border-top: 1px #ddd solid;
    border-bottom: 1px #ddd solid;
}

.showRawExceptionContainer {
    margin-top: 10px;
    margin-bottom: 10px;
}

.expandCollapseButton {
    cursor: pointer;
    float: left;
    height: 16px;
    width: 16px;
    font-size: 10px;
    position: absolute;
    left: 10px;
    background-color: #eee;
    padding: 0;
    border: 0;
    margin: 0;
}

        </style><style class="darkreader darkreader--sync" media="screen"></style>
    <style class="darkreader darkreader--override" media="screen">.jfk-bubble {
    background-color: #000000 !important;
}
.vimvixen-hint {
    background-color: #7b5300 !important;
    border-color: #d8b013 !important;
    color: #f3e8c8 !important;
}</style></head>
    <body>
        <h1>An unhandled panic occurred while processing the request.</h1>
            <div class="titleerror">Panic: {{ .Exception }}</div>
        <ul id="header">
            <li id="stack" tabindex="1" class="selected">
                Stack
            </li>
            <li id="query" tabindex="2" class="">
                Query
            </li>
            <li id="cookies" tabindex="3" class="">
                Cookies
            </li>
            <li id="headers" tabindex="4" class="">
                Headers
            </li>
        </ul>

        <div id="stackpage" class="page" style="">
            <ul>
                                    <li>
                        <h2 class="stackerror">Panic: {{ .Exception }}</h2>
                        <ul>
                            <li class="frame" id="frame1">
                                     <pre class="rawExceptionStackTrace">{{.StackTrace}}</pre>
                            </li>
                        </ul>
                    </li>
            </ul>
        </div>

        <div id="querypage" class="page" style="display: none;">
                <p>{{ .Query }}</p>
        </div>

        <div id="cookiespage" class="page" style="display: none;">
				{{ if .WithoutCookies }}
 					<p>No cookies data.</p>
				{{ end }} 
				{{ if .HaveCookies }}
                <table>
                    <thead>
                        <tr>
                            <th>Variable</th>
                            <th>Value</th>
                        </tr>
                    </thead>
                    <tbody>
							{{ range $key, $value := .Cookies }}
                            <tr>
                                <td>{{ $value.Name }}</td>
                                <td>{{ $value.Value }}</td>
                            </tr>
							{{ end }}
                    </tbody>
                </table>
				{{ end }}
        </div>
        <div id="headerspage" class="page" style="display: none;">
                <table>
                    <thead>
                        <tr>
                            <th>Variable</th>
                            <th>Value</th>
                        </tr>
                    </thead>
                    <tbody>
{{ range $key, $value := .Headers }}
<tr>
<td>{{$key}}</td>
<td>{{$value}}</td>
</tr>
{{ end }}
                    </tbody>
                </table>
        </div>
        <script>
            //<!--
            (function (window, undefined) {
    "use strict";

    function ns(selector, element) {
        return new NodeCollection(selector, element);
    }

    function NodeCollection(selector, element) {
        this.items = [];
        element = element || window.document;

        var nodeList;

        if (typeof (selector) === "string") {
            nodeList = element.querySelectorAll(selector);
            for (var i = 0, l = nodeList.length; i < l; i++) {
                this.items.push(nodeList.item(i));
            }
        }
    }

    NodeCollection.prototype = {
        each: function (callback) {
            for (var i = 0, l = this.items.length; i < l; i++) {
                callback(this.items[i], i);
            }
            return this;
        },

        children: function (selector) {
            var children = [];

            this.each(function (el) {
                children = children.concat(ns(selector, el).items);
            });

            return ns(children);
        },

        hide: function () {
            this.each(function (el) {
                el.style.display = "none";
            });

            return this;
        },

        toggle: function () {
            this.each(function (el) {
                el.style.display = el.style.display === "none" ? "" : "none";
            });

            return this;
        },

        show: function () {
            this.each(function (el) {
                el.style.display = "";
            });

            return this;
        },

        addClass: function (className) {
            this.each(function (el) {
                var existingClassName = el.className,
                    classNames;
                if (!existingClassName) {
                    el.className = className;
                } else {
                    classNames = existingClassName.split(" ");
                    if (classNames.indexOf(className) < 0) {
                        el.className = existingClassName + " " + className;
                    }
                }
            });

            return this;
        },

        removeClass: function (className) {
            this.each(function (el) {
                var existingClassName = el.className,
                    classNames, index;
                if (existingClassName === className) {
                    el.className = "";
                } else if (existingClassName) {
                    classNames = existingClassName.split(" ");
                    index = classNames.indexOf(className);
                    if (index > 0) {
                        classNames.splice(index, 1);
                        el.className = classNames.join(" ");
                    }
                }
            });

            return this;
        },

        attr: function (name) {
            if (this.items.length === 0) {
                return null;
            }

            return this.items[0].getAttribute(name);
        },

        on: function (eventName, handler) {
            this.each(function (el, idx) {
                var callback = function (e) {
                    e = e || window.event;
                    if (!e.which && e.keyCode) {
                        e.which = e.keyCode; // Normalize IE8 key events
                    }
                    handler.apply(el, [e]);
                };

                if (el.addEventListener) { // DOM Events
                    el.addEventListener(eventName, callback, false);
                } else if (el.attachEvent) { // IE8 events
                    el.attachEvent("on" + eventName, callback);
                } else {
                    el["on" + type] = callback;
                }
            });

            return this;
        },

        click: function (handler) {
            return this.on("click", handler);
        },

        keypress: function (handler) {
            return this.on("keypress", handler);
        }
    };

    function frame(el) {
        ns(".source .collapsible", el).toggle();
    }

    function expandCollapseButton(el) {
        var frameId = el.getAttribute("data-frameId");
        frame(document.getElementById(frameId));
        if (el.innerText === "+") {
            el.innerText = "-";
        }
        else {
            el.innerText = "+";
        }
    }

    function tab(el) {
        var unselected = ns("#header .selected").removeClass("selected").attr("id");
        var selected = ns("#" + el.id).addClass("selected").attr("id");

        ns("#" + unselected + "page").hide();
        ns("#" + selected + "page").show();
    }

    ns(".rawExceptionDetails").hide();
    ns(".collapsible").hide();
    ns(".page").hide();
    ns("#stackpage").show();

    ns(".expandCollapseButton")
        .click(function () {
            expandCollapseButton(this);
        })
        .keypress(function (e) {
            if (e.which === 13) {
                expandCollapseButton(this);
            }
        });

    ns("#header li")
        .click(function () {
            tab(this);
        })
        .keypress(function (e) {
            if (e.which === 13) {
                tab(this);
            }
        });

    ns(".showRawException")
        .click(function () {
            var exceptionDetailId = this.getAttribute("data-exceptionDetailId");
            ns("#" + exceptionDetailId).toggle();
        });
})(window);
            //-->
        </script>
    

</body></html>`
)
