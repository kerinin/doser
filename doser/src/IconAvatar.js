import React from 'react';

import { makeStyles } from '@material-ui/core/styles';
import Avatar from '@material-ui/core/Avatar';

const useStyles = makeStyles((theme) => ({
    avatar: {
        color: theme.palette.background.paper,
        backgroundColor: theme.palette.secondary.main,
    },
}));

function IconAvatar({ children }) {
    const classes = useStyles();
    return (
        <Avatar className={classes.avatar}>{children}</Avatar>
    )
}

export default IconAvatar;
