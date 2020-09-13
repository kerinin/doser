import React from 'react';
import CssBaseline from '@material-ui/core/CssBaseline';
import { createMuiTheme, makeStyles, ThemeProvider } from '@material-ui/core/styles';
import purple from '@material-ui/core/colors/purple';
import green from '@material-ui/core/colors/green';
import { VictoryBar } from 'victory';
import Button from '@material-ui/core/Button';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import IconButton from '@material-ui/core/IconButton';
import MenuIcon from '@material-ui/icons/Menu';
import Paper from '@material-ui/core/Paper';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import CardHeader from '@material-ui/core/CardHeader';
import VerticalAlignTopIcon from '@material-ui/icons/VerticalAlignTop';
import TimerIcon from '@material-ui/icons/Timer';
import Avatar from '@material-ui/core/Avatar';
import AutorenewIcon from '@material-ui/icons/Autorenew';

const theme = createMuiTheme({
  palette: {
    primary: {
      main: purple[500],
    },
    secondary: {
      main: green[500],
    },
  },
});

const useStyles = makeStyles({
  card: {
    margin: 20,
  },
});

function App() {
  const classes = useStyles();

  return (
    <React.Fragment>
      <CssBaseline />
      <ThemeProvider theme={theme}>
        <AppBar position="static">
          <Toolbar>
            <Typography variant="h6">
              Doser
            </Typography>
          </Toolbar>
        </AppBar>

        <Card className={classes.card}>
          <CardHeader avatar={<Avatar><VerticalAlignTopIcon /></Avatar>} title="Auto Top-Off" />
          <CardContent>
            <Button variant="contained" color="primary">Add ATO</Button>
          </CardContent>
        </Card>

        <Card className={classes.card}>
          <CardHeader avatar={<Avatar><AutorenewIcon /></Avatar>} title="Auto Water Change" />
          <CardContent>
            <Button variant="contained" color="primary">Add AWC</Button>
          </CardContent>
        </Card>

        <Card className={classes.card}>
          <CardHeader avatar={<Avatar><TimerIcon /></Avatar>} title="Doser" />
          <CardContent>
            <Button variant="contained" color="primary">Add Doser</Button>
          </CardContent>
        </Card>
      </ThemeProvider>
    </React.Fragment>
  );
}

export default App;
