import React from 'react';
import { createContext, useContext } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';
import Button from '@material-ui/core/Button';
import IconButton from '@material-ui/core/IconButton';
import FormHelperText from '@material-ui/core/FormHelperText';
import Menu from '@material-ui/core/Menu';
import MenuItem from '@material-ui/core/MenuItem';
import Input from '@material-ui/core/Input';
import InputLabel from '@material-ui/core/InputLabel';
import Chip from '@material-ui/core/Chip';
import Card from '@material-ui/core/Card';
import { CardActions, CardContent, CardHeader } from '@material-ui/core';
import { useQuery } from 'graphql-hooks'
import { Avatar, LinearProgress, ListItemAvatar, Typography } from '@material-ui/core';

import AddCircleIcon from '@material-ui/icons/AddCircle';

const QUERY = `query {
    water_level_sensors {
        id
        kind
    }
}`

const useStyles = makeStyles((theme) => ({
    grid: {
        flexGrow: 1,
    },
}));

export const Api = createContext(null);

function Content({ sensors }) {
    const [anchorEl, setAnchorEl] = React.useState(null);
    const { loading, error, data } = useQuery(QUERY, {})
    const classes = useStyles();
    const api = useContext(Api);

    const handleClick = (event) => {
        setAnchorEl(event.currentTarget);
    };

    const handleSelectSensor = (id) => {
        api.addSensor(id);
        setAnchorEl(null);
    };

    const handleClose = (id) => {
        setAnchorEl(null);
    };

    const handleDelete = (id) => {
        api.removeSensor(id)
    };

    if (loading) return (
        <CardContent><LinearProgress /></CardContent>
    )
    if (error) return (
        <CardContent>Failed to load water level sensors</CardContent>
    )
    if (data.water_level_sensors == null) return (
        <CardContent>No water level sensors defined - create one first</CardContent>
    )

    return (
        <CardContent>
            <Grid container spacing={2} className={classes.grid}>
                <Grid item xs={4}>
                    <InputLabel>Sensors</InputLabel>
                    {sensors.map((sensor, idx) => <Chip key={idx} label={sensor} onDelete={() => handleDelete(sensor)} />)}
                    <IconButton onClick={handleClick} ><AddCircleIcon /></IconButton>
                    <Menu
                        id="add-sensor-menu"
                        anchorEl={anchorEl}
                        keepMounted
                        open={Boolean(anchorEl)}
                        onClose={handleClose}
                    >
                        {data.water_level_sensors.map((s) => <MenuItem key={s.id} onClick={(e) => handleSelectSensor(s.id)}>{s.id}</MenuItem>)}
                    </Menu>
                </Grid>
                <Grid item xs={4}>
                    <InputLabel>Fill Rate</InputLabel>
                    <Input id="input-rate" />
                    <FormHelperText>Rate in mL/min</FormHelperText>
                </Grid>
                <Grid item xs={4}>
                    <InputLabel>Fill Interval</InputLabel>
                    <Input id="input-interval" />
                    <FormHelperText>Interval in minutes</FormHelperText>
                </Grid>
            </Grid>
        </CardContent>
    )
}

function EditAutoTopOff({ sensors }) {
    return (
        <Card>
            <CardHeader title="Settings" />
            <Content sensors={sensors} addSensor />
            <CardActions>
                <Button color="primary">Cancel</Button>
                <Button variant="contained" color="primary">Save</Button>
            </CardActions>
        </Card>
    )
}

export default EditAutoTopOff;