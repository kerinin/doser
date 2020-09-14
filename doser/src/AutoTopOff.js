import React from 'react';
import { useParams } from "react-router-dom";
import { makeStyles } from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';
import Box from '@material-ui/core/Box';
import Paper from '@material-ui/core/Paper';
import { VictoryBar } from 'victory';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import TablePagination from '@material-ui/core/TablePagination';

import DoserAppBar from './DoserAppBar';
import EditAutoTopOff from './EditAutoTopOff';

const useStyles = makeStyles((theme) => ({
    grid: {
        flexGrow: 1,
    },
}));

function AutoTopOff() {
    const classes = useStyles();
    let { atoId } = useParams();

    const handleClick = () => {
        console.info('You clicked the Chip.');
    };

    const handleDelete = () => {
        console.info('You clicked the delete icon.');
    };

    return (
        <React.Fragment>
            <DoserAppBar />
            <Box m={2}>
                <EditAutoTopOff />

                <Grid container spacing={2} className={classes.grid}>
                    <Grid item xs={12}>
                        <Paper>

                            <VictoryBar />

                            <TableContainer>
                                <Table>
                                    <TableHead>
                                        <TableRow>
                                            <TableCell>Time</TableCell>
                                            <TableCell>Dose</TableCell>
                                        </TableRow>
                                    </TableHead>
                                    <TableBody>
                                        <TableRow>
                                            <TableCell>Today</TableCell>
                                            <TableCell>124mL</TableCell>
                                        </TableRow>
                                    </TableBody>
                                    {/* <TablePagination
                                            rowsPerPageOptions={[10, 50]}
                                            count={30}
                                            rowsPerPage={10}
                                            page={1}
                                        /> */}
                                </Table>
                            </TableContainer>
                        </Paper>
                    </Grid>
                </Grid>
            </Box>
        </React.Fragment >
    )
}

export default AutoTopOff;