import React from 'react';
import { useParams } from "react-router-dom";
import { makeStyles } from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';
import Box from '@material-ui/core/Box';
import Paper from '@material-ui/core/Paper';
import { VictoryArea, VictoryAxis, VictoryBar, VictoryChart, VictoryLine, VictoryTheme } from 'victory';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import TablePagination from '@material-ui/core/TablePagination';
import { useTheme } from '@material-ui/core/styles';

import DoserAppBar from './DoserAppBar';
import EditAutoTopOff from './EditAutoTopOff';
import { Card, CardContent, CardHeader, CardMedia } from '@material-ui/core';

const useStyles = makeStyles((theme) => ({
    grid: {
        flexGrow: 1,
    },
    line: {
        stroke: theme.palette.primary.main,
    },
}));

function AutoTopOff() {
    const theme = useTheme();
    const classes = useStyles();
    let { atoId } = useParams();

    const victoryTheme = {
        ...VictoryTheme.material,
        chart: {
            ...VictoryTheme.material.chart,
            width: 800,
            height: 400,
        },

        line: {
            ...VictoryTheme.material.line,
            style: {
                ...VictoryTheme.material.line.style,
                data: {
                    ...VictoryTheme.material.line.style.data,
                    stroke: theme.palette.primary.main,
                    strokeWidth: 4,
                },
            },
        },
    }

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

                <Grid container spacing={2} className={classes.grid}>
                    <Grid item xs={12}>
                        <EditAutoTopOff />
                    </Grid>

                    <Grid item xs={12}>
                        <Card>
                            <CardHeader title="History" />
                            <CardContent>
                                <VictoryChart
                                    theme={victoryTheme}
                                >
                                    <VictoryArea
                                        interpolation="stepAfter"
                                        minDomain={{ y: 0 }}
                                        data={
                                            [
                                                { x: 1, y: 19000 },
                                                { x: 2, y: 0 },
                                                { x: 3, y: 16500 },
                                                { x: 4, y: 0 },
                                            ]
                                        }
                                        style={{
                                            data: {
                                                fill: theme.palette.secondary.main,
                                            }
                                        }}
                                    />
                                    <VictoryLine
                                        interpolation="natural"
                                        minDomain={{ y: 0 }}
                                        data={
                                            [
                                                { x: 1, y: 13000 },
                                                { x: 2, y: 16500 },
                                                { x: 3, y: 14250 },
                                                { x: 4, y: 19000 }
                                            ]
                                        }

                                    />
                                    <VictoryAxis></VictoryAxis>
                                    <VictoryAxis dependentAxis></VictoryAxis>
                                </VictoryChart>
                            </CardContent>

                            <CardContent>
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
                            </CardContent>
                        </Card>
                    </Grid>
                </Grid>
            </Box>
        </React.Fragment >
    )
}

export default AutoTopOff;