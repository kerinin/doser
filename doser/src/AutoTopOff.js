import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import {
  VictoryArea,
  VictoryAxis,
  VictoryChart,
  VictoryLine,
  VictoryTheme,
} from "victory";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableContainer from "@material-ui/core/TableContainer";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import { useTheme } from "@material-ui/core/styles";
import { useQuery } from "graphql-hooks";
import { useMutation } from "graphql-hooks";

import DoserAppBar from "./DoserAppBar";
import EditAutoTopOff from "./EditAutoTopOff";
import { Api as EditAutoTopOffApi } from "./EditAutoTopOff";
import { Card, CardContent, CardHeader, Typography } from "@material-ui/core";
import { useParams } from "react-router-dom";

const useStyles = makeStyles((theme) => ({
  grid: {
    flexGrow: 1,
  },
  line: {
    stroke: theme.palette.primary.main,
  },
}));

const QUERY = `query GetAutoTopOff($id: ID!) {
    auto_top_off(id: $id) {
        id
        pump {
          id
        }
        level_sensors {
          id
        }
        fill_rate
        fill_interval
        max_fill_volume
        events {
          id
          timestamp
          kind
          data
        }
    }
}`;

const EDIT = `mutation EditAutoTopOff($id: ID!, $pump_id: ID!, $level_sensors: [ID!]!, $fill_rate: Float!, $fill_interval: Int!, $max_fill_volume: Float!) {
    updateAutoTopOff(
        id: $id,
        pump_id: $pump_id,
        level_sensors: $level_sensors,
        fill_rate: $fill_rate,
        fill_interval: $fill_interval,
        max_fill_volume: $max_fill_volume,
    ) {
        id
    }
}`;

function AutoTopOff({ id }) {
  const { atoId } = useParams();
  const { loading, error, data, refetch } = useQuery(QUERY, {
    variables: { id: atoId },
  });
  const classes = useStyles();

  if (loading)
    return (
      <React.Fragment>
        <DoserAppBar />
      </React.Fragment>
    );
  if (error)
    return (
      <Grid item xs={12}>
        <Typography>Failed to load ATO</Typography> />
      </Grid>
    );

  if (!data.auto_top_off)
    return (
      <React.Fragment>
        <DoserAppBar />
        <Box m={2}>
          <Typography variant="h6">Not Found</Typography>
          <Typography variant="body">
            '{atoId}' is not a recognized ATO
          </Typography>
        </Box>
      </React.Fragment>
    );

  return (
    <React.Fragment>
      <DoserAppBar />
      <Box m={2}>
        <Grid container spacing={2} className={classes.grid}>
          <Content ato={data.auto_top_off} reload={refetch} />
        </Grid>
      </Box>
    </React.Fragment>
  );
}

function Content({ ato, reload }) {
  const [editAutoTopOff, { error }] = useMutation(EDIT);

  // TODO: Fill with actual data!
  const [pump, setPump] = React.useState(ato.pump.id);
  const [sensors, setSensors] = React.useState(
    ato.level_sensors.map((s) => s.id)
  );
  const [fillRate, setFillRate] = React.useState(ato.fill_rate);
  const [fillInterval, setFillInterval] = React.useState(ato.fill_interval);
  const [maxFillVolume, setMaxFillVolume] = React.useState(ato.max_fill_volume);

  function addSensor(id) {
    setSensors(sensors.concat(id));
  }
  function removeSensor(id) {
    setSensors(sensors.filter((item) => item !== id));
  }
  function cancel() {
    setPump(ato.pump.id);
    setSensors(ato.level_sensors.map((s) => s.id));
    setFillRate(ato.fill_rate);
    setFillInterval(ato.fill_interval);
    setMaxFillVolume(ato.max_fill_volume);
  }
  function submit() {
    editAutoTopOff({
      variables: {
        id: ato.id,
        pump_id: pump,
        level_sensors: sensors,
        fill_rate: fillRate,
        fill_interval: fillInterval,
        max_fill_volume: maxFillVolume,
      },
    }).then(({ error }) => {
      if (!error) reload();
    });
  }

  const api = {
    pump,
    setPump,
    sensors,
    addSensor,
    removeSensor,
    fillRate,
    setFillRate,
    fillInterval,
    setFillInterval,
    maxFillVolume,
    setMaxFillVolume,
    cancel,
    submit,
    error,
  };

  const theme = useTheme();

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
  };

  return (
    <React.Fragment>
      <Grid item xs={12}>
        <EditAutoTopOffApi.Provider value={api}>
          <EditAutoTopOff />
        </EditAutoTopOffApi.Provider>
      </Grid>

      <Grid item xs={12}>
        <Card>
          <CardHeader title="History" />
          <CardContent>
            <VictoryChart theme={victoryTheme}>
              <VictoryArea
                interpolation="stepAfter"
                minDomain={{ y: 0 }}
                data={[
                  { x: 1, y: 19000 },
                  { x: 2, y: 0 },
                  { x: 3, y: 16500 },
                  { x: 4, y: 0 },
                ]}
                style={{
                  data: {
                    fill: theme.palette.secondary.main,
                  },
                }}
              />
              <VictoryLine
                interpolation="natural"
                minDomain={{ y: 0 }}
                data={[
                  { x: 1, y: 13000 },
                  { x: 2, y: 16500 },
                  { x: 3, y: 14250 },
                  { x: 4, y: 19000 },
                ]}
              />
              <VictoryAxis></VictoryAxis>
              <VictoryAxis dependentAxis></VictoryAxis>
            </VictoryChart>
          </CardContent>

          <CardContent>
            <EventsTable ato={ato} />
          </CardContent>
        </Card>
      </Grid>
    </React.Fragment>
  );
}

function EventsTable({ ato }) {
  return (
    <TableContainer>
      <Table>
        <TableHead>
          <TableRow>
            <TableCell>Time</TableCell>
            <TableCell>Event</TableCell>
            <TableCell>Message</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          <EventsTableRows ato={ato} />
        </TableBody>
      </Table>
    </TableContainer>
  );
}

function EventsTableRows({ ato }) {
  if (!ato.events)
    return (
      <TableRow>
        <TableCell colSpan={3}>No Events</TableCell>
      </TableRow>
    );

  return ato.events.map(({ timestamp, kind, data }) => (
    <TableRow>
      <TableCell>{timestamp}</TableCell>
      <TableCell>{kind}</TableCell>
      <TableCell>{data}</TableCell>
    </TableRow>
  ));
}

export default AutoTopOff;
