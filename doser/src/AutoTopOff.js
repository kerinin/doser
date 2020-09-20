import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import {
  Point,
  VictoryArea,
  VictoryAxis,
  VictoryChart,
  VictoryLabel,
  VictoryLine,
  VictoryPortal,
  VictoryScatter,
  VictoryTheme,
  VictoryTooltip,
  VictoryVoronoiContainer,
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
import { DataGrid } from "@material-ui/data-grid";

import DoserAppBar from "./DoserAppBar";
import EditAutoTopOff from "./EditAutoTopOff";
import { Api as EditAutoTopOffApi } from "./EditAutoTopOff";
import {
  ButtonGroup,
  Button,
  Card,
  CardContent,
  CardHeader,
  Typography,
} from "@material-ui/core";
import { useParams } from "react-router-dom";

const useStyles = makeStyles((theme) => ({
  grid: {
    flexGrow: 1,
  },
  line: {
    stroke: theme.palette.primary.main,
  },
}));

function useVictoryTheme() {
  const theme = useTheme();
  return {
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

    scatter: {
      ...VictoryTheme.material.scatter,
      style: {
        ...VictoryTheme.material.scatter.style,
        data: {
          ...VictoryTheme.material.scatter.style.data,
          fill: theme.palette.secondary.main,
        },
      },
    },
  };
}

const QUERY = `query GetAutoTopOff($id: ID!) {
    auto_top_off(id: $id) {
        id
        pump {
          id
          history {
            timestamp
            volume
          }
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
        rate(window: 21600) {
          timestamp
          rate
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
            <Chart ato={ato} />
          </CardContent>
        </Card>
      </Grid>
    </React.Fragment>
  );
}

function Chart({ ato }) {
  const victoryTheme = useVictoryTheme();

  return (
    <VictoryChart
      theme={victoryTheme}
      minDomain={{ y: 0 }}
      domainPadding={{ y: [20, 20] }}
      containerComponent={<VictoryVoronoiContainer />}
    >
      <VictoryScatter
        labelComponent={<VictoryTooltip />}
        data={ato.pump.history.map(({ timestamp, volume }) => ({
          x: (timestamp - Date.now() / 1000) / 60 / 60 / 24,
          y: volume,
          label: `${new Date(timestamp * 1000).toISOString()}: ${volume}mL`,
        }))}
      />
      <VictoryLine
        interpolation="bundle"
        data={ato.rate.map(({ timestamp, rate }) => ({
          x: (timestamp - Date.now() / 1000) / 60 / 60 / 24,
          y: rate,
        }))}
      />
      <VictoryAxis
        label="days ago"
        axisLabelComponent={<VictoryLabel dy={20} />}
      />
      <VictoryAxis
        dependentAxis
        label="Rate (mL/h)"
        axisLabelComponent={<VictoryLabel dy={30} />}
      />
    </VictoryChart>
  );
}

function EventsTable({ ato }) {
  const columns = [
    {
      field: "timestamp",
      headerName: "Time",
      width: 250,
      type: "datetime",
      valueFormatter: ({ value }) => new Date(value * 1000).toISOString(),
    },
    { field: "volume", headerName: "Volume", width: 100 },
  ];

  return (
    <React.Fragment>
      <div style={{ height: 400, width: "100%" }}>
        <DataGrid rows={ato.pump.history} columns={columns} pageSize={10} />
      </div>
    </React.Fragment>
  );
}

export default AutoTopOff;
