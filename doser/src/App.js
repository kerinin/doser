import React from 'react';
import CssBaseline from '@material-ui/core/CssBaseline';
import { createMuiTheme, ThemeProvider } from '@material-ui/core/styles';
import { makeStyles } from '@material-ui/core/styles';
import purple from '@material-ui/core/colors/purple';
import yellow from '@material-ui/core/colors/yellow';
import Grid from '@material-ui/core/Grid';
import Box from '@material-ui/core/Box';

import DoserAppBar from './DoserAppBar';
import AutoTopOffs from './AutoTopOffs';
import AutoWaterChanges from './AutoWaterChanges';
import Dosers from './Dosers';


const theme = createMuiTheme({
  palette: {
    primary: {
      main: purple[500],
    },
    secondary: {
      main: yellow[500],
    },
  },
});

const useStyles = makeStyles((theme) => ({
  grid: {
    flexGrow: 1,
  },
}));

function App() {
  const classes = useStyles();

  return (
    <React.Fragment>
      <CssBaseline />
      <ThemeProvider theme={theme}>
        <DoserAppBar />
        <Box m={2}>
          <Grid container spacing={2} className={classes.grid}>
            <Grid item xs={6}>
              <AutoTopOffs />
            </Grid>
            <Grid item xs={6}>
              <AutoWaterChanges />
            </Grid>
            <Grid item xs={6}>
              <Dosers />
            </Grid>
          </Grid>
        </Box>
      </ThemeProvider>
    </React.Fragment>
  );
}

export default App;
