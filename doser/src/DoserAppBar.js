import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Toolbar from '@material-ui/core/Toolbar';
import AppBar from '@material-ui/core/AppBar';
import Typography from '@material-ui/core/Typography';
import Switch from '@material-ui/core/Switch';
import IconButton from '@material-ui/core/IconButton';

import WavesIcon from '@material-ui/icons/Waves';
import SettingsIcon from '@material-ui/icons/Settings';

const useStyles = makeStyles({
    root: {
        flexGrow: 1,
    },
    title: {
        flexGrow: 1,
    }
});

function DoserAppBar() {
    const classes = useStyles();

    return (
        <AppBar position="static" className={classes.root}>
            <Toolbar>
                <WavesIcon />
                <Typography variant="h6" className={classes.title}>Dozer</Typography>

                <Switch />
                <IconButton><SettingsIcon /></IconButton>
            </Toolbar>
        </AppBar>
    )
}

export default DoserAppBar;