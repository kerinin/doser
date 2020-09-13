import React from 'react';

import Button from '@material-ui/core/Button';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import CardHeader from '@material-ui/core/CardHeader';
import CardActions from '@material-ui/core/CardActions';

import AutorenewIcon from '@material-ui/icons/Autorenew';
import AddCircleIcon from '@material-ui/icons/AddCircle';

import IconAvatar from './IconAvatar';

function AutoWaterChanges() {
    return (
        <Card>
            <CardHeader avatar={<IconAvatar><AutorenewIcon /></IconAvatar>} title="Auto Water Change" />
            <CardContent> </CardContent>
            <CardActions>
                <Button startIcon={<AddCircleIcon />}>Add AWC</Button>
            </CardActions>
        </Card>
    )
}

export default AutoWaterChanges;
