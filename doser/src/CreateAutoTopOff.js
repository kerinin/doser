import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';
import Box from '@material-ui/core/Box';

import DoserAppBar from './DoserAppBar';
import { Api as EditAutoTopOffApi } from './EditAutoTopOff';
import EditAutoTopOff from './EditAutoTopOff';

const useStyles = makeStyles((theme) => ({
    grid: {
        flexGrow: 1,
    },
    line: {
        stroke: theme.palette.primary.main,
    },
}));

function CreateAutoTopOff() {
    const classes = useStyles();
    const [sensors, setSensors] = React.useState([]);

    function addSensor(id) {
        setSensors(sensors.concat(id))
    }
    function removeSensor(id) {
        setSensors(sensors.filter(item => item != id))
    }

    const api = { addSensor, removeSensor };
    return (
        <EditAutoTopOffApi.Provider value={api}>
            <DoserAppBar />
            <Box m={2}>

                <Grid container spacing={2} className={classes.grid}>
                    <Grid item xs={12}>
                        <EditAutoTopOff sensors={sensors} />
                    </Grid>
                </Grid>
            </Box>
        </EditAutoTopOffApi.Provider>
    )
}

export default CreateAutoTopOff;
