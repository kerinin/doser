import React from "react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import CssBaseline from "@material-ui/core/CssBaseline";
import {
  unstable_createMuiStrictModeTheme,
  ThemeProvider,
} from "@material-ui/core/styles";
import purple from "@material-ui/core/colors/purple";
import yellow from "@material-ui/core/colors/yellow";
import { GraphQLClient, ClientContext } from "graphql-hooks";

import Dashboard from "./Dashboard";
import AutoTopOff from "./AutoTopOff";
import AutoWaterChange from "./AutoWaterChange";
import CreateAutoTopOff from "./CreateAutoTopOff";
import CreateAutoWaterChange from "./CreateAutoWaterChange";

const theme = unstable_createMuiStrictModeTheme({
  palette: {
    primary: {
      main: purple[500],
    },
    secondary: {
      main: yellow[500],
    },
  },
});

const client = new GraphQLClient({
  url: "http://raspberrypi.local:8080/query",
});

function App() {
  return (
    <ClientContext.Provider value={client}>
      <CssBaseline />
      <ThemeProvider theme={theme}>
        <Router>
          <Switch>
            <Route path="/ato/create">
              <CreateAutoTopOff />
            </Route>
            <Route path="/ato/:atoId">
              <AutoTopOff />
            </Route>
            <Route path="/awc/create">
              <CreateAutoWaterChange />
            </Route>
            <Route path="/awc/:awcId">
              <AutoWaterChange />
            </Route>
            <Route path="/">
              <Dashboard />
            </Route>
          </Switch>
        </Router>
      </ThemeProvider>
    </ClientContext.Provider>
  );
}

export default App;
