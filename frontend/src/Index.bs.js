'use strict';

var React = require("react");
var App$Doser = require("./App.bs.js");
var ReactDOMRe = require("reason-react/src/legacy/ReactDOMRe.bs.js");

ReactDOMRe.renderToElementWithId(React.createElement(App$Doser.make, {}), "root");

/*  Not a pure module */
