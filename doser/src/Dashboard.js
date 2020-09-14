import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';
import Box from '@material-ui/core/Box';

import AutoTopOffs from './AutoTopOffs';
import AutoWaterChanges from './AutoWaterChanges';
import Dosers from './Dosers';
import DoserAppBar from './DoserAppBar';


const useStyles = makeStyles((theme) => ({
    grid: {
        flexGrow: 1,
    },
}));

function Dashboard() {
    const classes = useStyles();

    return (
        <React.Fragment>
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
        </React.Fragment>
    )
}

export default Dashboard;