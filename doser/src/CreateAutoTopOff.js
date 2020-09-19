import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import { useMutation } from "graphql-hooks";
import { useHistory } from "react-router-dom";

import DoserAppBar from "./DoserAppBar";
import { Api as EditAutoTopOffApi } from "./EditAutoTopOff";
import EditAutoTopOff from "./EditAutoTopOff";

const CREATE = `mutation CreateAutoTopOff($pump_id: ID!, $sensors: [ID!]!, $fill_rate: Float!, $fill_interval: Int!, $max_fill_volume: Float!) {
    createAutoTopOff(
        pump_id: $pump_id,
        level_sensors: $sensors,
        fill_rate: $fill_rate,
        fill_interval: $fill_interval,
        max_fill_volume: $max_fill_volume,
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

function CreateAutoTopOff() {
  let history = useHistory();

  const classes = useStyles();
  const [pump, setPump] = React.useState("");
  const [sensors, setSensors] = React.useState([]);
  const [fillRate, setFillRate] = React.useState(100);
  const [fillInterval, setFillInterval] = React.useState(10);
  const [maxFillVolume, setMaxFillVolume] = React.useState(1000);

  const [createAutoTopOff, { loading, error }] = useMutation(CREATE);

  function addSensor(id) {
    setSensors(sensors.concat(id));
  }
  function removeSensor(id) {
    setSensors(sensors.filter((item) => item !== id));
  }
  function submit() {
    createAutoTopOff({
      variables: {
        pump_id: pump,
        sensors: sensors,
        fill_rate: fillRate,
        fill_interval: fillInterval,
        max_fill_volume: maxFillVolume,
      },
    }).then(({ data, error }) => {
      if (!error) history.push(`/ato/${data.createAutoTopOff.id}`);
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
    submit,
    loading,
    error,
  };

  return (
    <EditAutoTopOffApi.Provider value={api}>
      <DoserAppBar />
      <Box m={2}>
        <Grid container spacing={2} className={classes.grid}>
          <Grid item xs={12}>
            <EditAutoTopOff sensors={sensors} />
          </Grid>
        </Grid>
      </Box>
    </EditAutoTopOffApi.Provider>
  );
}

export default CreateAutoTopOff;
