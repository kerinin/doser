import React from 'react';
import Button from '@material-ui/core/Button';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import CardHeader from '@material-ui/core/CardHeader';
import CardActions from '@material-ui/core/CardActions';

import TimerIcon from '@material-ui/icons/Timer';
import AddCircleIcon from '@material-ui/icons/AddCircle';

import IconAvatar from './IconAvatar';

function Dosers() {
    return (
        <Card>
            <CardHeader avatar={<IconAvatar><TimerIcon /></IconAvatar>} title="Doser" />
            <CardContent> </CardContent>
            <CardActions>
                <Button startIcon={<AddCircleIcon />}>Add Doser</Button>
            </CardActions>
        </Card>
    )
}

export default Dosers;