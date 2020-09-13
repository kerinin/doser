import React from 'react';

import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import CardActions from '@material-ui/core/CardActions';
import CardHeader from '@material-ui/core/CardHeader';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemSecondaryAction from '@material-ui/core/ListItemSecondaryAction';
import ListItemText from '@material-ui/core/ListItemText';
import Button from '@material-ui/core/Button';
import Switch from '@material-ui/core/Switch';
import IconButton from '@material-ui/core/IconButton';

import VerticalAlignTopIcon from '@material-ui/icons/VerticalAlignTop';
import DeleteIcon from '@material-ui/icons/Delete';
import AddCircleIcon from '@material-ui/icons/AddCircle';

import IconAvatar from './IconAvatar';


const useStyles = makeStyles((theme) => ({
    avatar: {
        color: theme.palette.background.paper,
        backgroundColor: theme.palette.primary.main,
    },
}));

function AutoTopOffs() {
    const classes = useStyles();

    return (
        <Card>
            <CardHeader avatar={<IconAvatar><VerticalAlignTopIcon /></IconAvatar>} title="Auto Top-Off" />
            <CardContent>
                <List>
                    <ListItem button>
                        <ListItemText primary={"ATO Name"} />
                        <ListItemSecondaryAction>
                            <Switch />
                            <IconButton edge="end">
                                <DeleteIcon />
                            </IconButton>
                        </ListItemSecondaryAction>
                    </ListItem>
                </List>
            </CardContent>
            <CardActions>
                <Button startIcon={<AddCircleIcon />}>Add ATO</Button>
            </CardActions>
        </Card >
    )
}

export default AutoTopOffs;