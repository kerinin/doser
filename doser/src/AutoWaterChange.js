import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import {
  createContainer,
  VictoryAxis,
  VictoryChart,
  VictoryLabel,
  VictoryLine,
  VictoryScatter,
  VictoryTheme,
  VictoryTooltip,
} from "victory";
import { useTheme } from "@material-ui/core/styles";
import { useQuery } from "graphql-hooks";
import { useMutation } from "graphql-hooks";

import DoserAppBar from "./DoserAppBar";
import EditAutoWaterChange from "./EditAutoWaterChange";
import { Api as EditAutoWaterChangeApi } from "./EditAutoWaterChange";
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

const QUERY = `query GetAutoWaterChange($id: ID!) {
    auto_water_change(id: $id) {
        id
        fresh_pump {
          id
        }
        waste_pump {
          id
        }
        exchange_rate
    }
}`;

const EDIT = `mutation EditAutoWaterChange($id: ID!, $fresh_pump_id: ID!, $waste_pump_id: ID!, $exchange_rate: Float!) {
    updateAutoWaterChange(
        id: $id,
        fresh_pump_id: $fresh_pump_id,
        waste_pump_id: $waste_pump_id,
        exchange_rate: $exchange_rate,
    ) {
        id
    }
}`;

function AutoWaterChange({ id }) {
  const { awcId } = useParams();
  const { loading, error, data, refetch } = useQuery(QUERY, {
    variables: { id: awcId },
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
        <Typography>Failed to load AWC</Typography> />
      </Grid>
    );

  if (!data.auto_top_off)
    return (
      <React.Fragment>
        <DoserAppBar />
        <Box m={2}>
          <Typography variant="h6">Not Found</Typography>
          <Typography variant="body">
            '{awcId}' is not a recognized AWC
          </Typography>
        </Box>
      </React.Fragment>
    );

  return (
    <React.Fragment>
      <DoserAppBar />
      <Box m={2}>
        <Grid container spacing={2} className={classes.grid}>
          <Content awc={data.auto_water_change} reload={refetch} />
        </Grid>
      </Box>
    </React.Fragment>
  );
}

function Content({ awc, reload }) {
  const [editAutoWaterChange, { error }] = useMutation(EDIT);

  const [freshPump, setFreshPump] = React.useState(awc.fresh_pump.id);
  const [wastePump, setWastePump] = React.useState(awc.waste_pump.id);
  const [exchangeRate, setExchangeRate] = React.useState(awc.exchange_rate);

  function cancel() {
    setFreshPump(awc.fresh_pump.id);
    setWastePump(awc.waste_pump.id);
    setExchangeRate(awc.exchange_rate);
  }
  function submit() {
    editAutoWaterChange({
      variables: {
        id: awc.id,
        fresh_pump_id: freshPump,
        waste_pump_id: wastePump,
        exchange_rate: exchangeRate,
      },
    }).then(({ error }) => {
      if (!error) reload();
    });
  }

  const api = {
    freshPump,
    setFreshPump,
    wastePump,
    setWastePump,
    exchangeRate,
    setExchangeRate,
    cancel,
    submit,
    error,
  };

  return (
    <React.Fragment>
      <Grid item xs={12} lg={6}>
        <EditAutoWaterChangeApi.Provider value={api}>
          <EditAutoWaterChange />
        </EditAutoWaterChangeApi.Provider>
      </Grid>
    </React.Fragment>
  );
}

export default AutoWaterChange;
