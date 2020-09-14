import React from 'react';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from "react-router-dom";
import CssBaseline from '@material-ui/core/CssBaseline';
import { unstable_createMuiStrictModeTheme, ThemeProvider } from '@material-ui/core/styles';
import purple from '@material-ui/core/colors/purple';
import yellow from '@material-ui/core/colors/yellow';

import Dashboard from './Dashboard';
import AutoTopOff from './AutoTopOff';

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

function App() {
  return (
    <React.Fragment>
      <CssBaseline />
      <ThemeProvider theme={theme}>
        <Router>
          <Switch>
            <Route path="/ato/:atoId">
              <AutoTopOff />
            </Route>
            <Route path="/">
              <Dashboard />
            </Route>
          </Switch>
        </Router>
      </ThemeProvider>
    </React.Fragment>
  );
}

export default App;
