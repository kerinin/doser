import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';
import Box from '@material-ui/core/Box';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import IconButton from '@material-ui/core/IconButton';
import FormHelperText from '@material-ui/core/FormHelperText';
import Menu from '@material-ui/core/Menu';
import MenuItem from '@material-ui/core/MenuItem';
import Input from '@material-ui/core/Input';
import InputLabel from '@material-ui/core/InputLabel';
import Chip from '@material-ui/core/Chip';
import Card from '@material-ui/core/Card';
import IconAvatar from './IconAvatar';
import { CardActions, CardContent, CardHeader } from '@material-ui/core';

import VerticalAlignTopIcon from '@material-ui/icons/VerticalAlignTop';
import AddCircleIcon from '@material-ui/icons/AddCircle';

const useStyles = makeStyles((theme) => ({
    grid: {
        flexGrow: 1,
    },
}));

function EditAutoTopOff() {
    const [anchorEl, setAnchorEl] = React.useState(null);

    const classes = useStyles();

    const handleClick = (event) => {
        setAnchorEl(event.currentTarget);
    };

    const handleClose = () => {
        setAnchorEl(null);
    };

    const handleDelete = () => {
        console.info('You clicked the delete icon.');
    };

    return (
        <Card>
            <CardHeader title="Settings" />
            <CardContent>
                <Grid container spacing={2} className={classes.grid}>
                    <Grid item xs={4}>
                        <InputLabel>Sensors</InputLabel>
                        <Chip label="Sensor 1" onDelete={handleDelete} />
                        <IconButton onClick={handleClick} ><AddCircleIcon /></IconButton>
                        <Menu
                            id="add-sensor-menu"
                            anchorEl={anchorEl}
                            keepMounted
                            open={Boolean(anchorEl)}
                            onClose={handleClose}
                        >
                            <MenuItem onClick={handleClose}>Sensor 1</MenuItem>
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
            <CardActions>
                <Button color="primary">Cancel</Button>
                <Button variant="contained" color="primary">Save</Button>
            </CardActions>
        </Card>
    )
}

export default EditAutoTopOff;