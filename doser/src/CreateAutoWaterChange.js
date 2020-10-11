import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import { useMutation } from "graphql-hooks";
import { useHistory } from "react-router-dom";

import DoserAppBar from "./DoserAppBar";
import { Api as EditAutoWaterChangeApi } from "./EditAutoWaterChange";
import EditAutoWaterChange from "./EditAutoWaterChange";

const CREATE = `mutation CreateAutoWaterChange($fresh_pump_id: ID!, $waste_pump_id: ID!, $exchange_rate: Float!) {
    createAutoWaterChange(
        fresh_pump_id: $fresh_pump_id,
        waste_pump_id: $waste_pump_id,
        exchange_rate: $exchange_rate
    ) {
        id
    }
}`;

const useStyles = makeStyles((theme) => ({
  grid: {
    flexGrow: 1,
  },
  line: {
    stroke: theme.palette.primary.main,
  },
}));

function CreateAutoWaterChange() {
  let history = useHistory();

  const classes = useStyles();
  const [freshPump, setFreshPump] = React.useState("");
  const [wastePump, setWastePump] = React.useState("");
  const [exchangeRate, setExchangeRate] = React.useState(1);

  const [createAutoWaterChange, { error }] = useMutation(CREATE);

  function cancel() {
    history.push("/");
  }
  function submit() {
    createAutoWaterChange({
      variables: {
        fresh_pump_id: freshPump,
        waste_pump_id: wastePump,
        exchange_rate: exchangeRate,
      },
    }).then(({ data, error }) => {
      if (!error) history.push(`/awc/${data.createAutoWaterChange.id}`);
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
    <EditAutoWaterChangeApi.Provider value={api}>
      <DoserAppBar />
      <Box m={2}>
        <Grid container spacing={2} className={classes.grid}>
          <Grid item xs={12}>
            <EditAutoWaterChange />
          </Grid>
        </Grid>
      </Box>
    </EditAutoWaterChangeApi.Provider>
  );
}

export default CreateAutoWaterChange;
