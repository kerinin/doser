import React from "react";
import { createContext, useContext } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Grid from "@material-ui/core/Grid";
import Button from "@material-ui/core/Button";
import FormControl from "@material-ui/core/FormControl";
import MenuItem from "@material-ui/core/MenuItem";
import Select from "@material-ui/core/Select";
import InputLabel from "@material-ui/core/InputLabel";
import Chip from "@material-ui/core/Chip";
import TextField from "@material-ui/core/TextField";
import Card from "@material-ui/core/Card";
import {
  CardActions,
  CardContent,
  CardHeader,
  Checkbox,
  FormControlLabel,
  FormGroup,
  FormLabel,
} from "@material-ui/core";
import { useQuery } from "graphql-hooks";
import { Typography } from "@material-ui/core";

const QUERY = `query {
    water_level_sensors {
        id
        kind
    }
    pumps {
        id
    }
}`;

const useStyles = makeStyles((theme) => ({
  grid: {
    flexGrow: 1,
  },
  form: {
    "& .MuiTextField-root": {
      margin: theme.spacing(1),
    },
  },
  formControl: {
    margin: theme.spacing(1),
    minWidth: 120,
  },
  chip: {
    margin: theme.spacing(1),
  },
}));

export const Api = createContext(null);

function Content() {
  const { loading, error, data } = useQuery(QUERY, {});
  const classes = useStyles();
  const api = useContext(Api);

  const handleSelectSensor = (id) => {
    api.addSensor(id);
  };

  if (loading) return <CardContent />;
  if (error)
    return <CardContent>Failed to load water level sensors</CardContent>;

  if (data.water_level_sensors == null)
    return (
      <CardContent>
        No water level sensors defined - create one first
      </CardContent>
    );

  return (
    <CardContent>
      <Grid container spacing={2} className={classes.grid}>
        <Grid item xs={12}>
          <FormControl className={classes.formControl}>
            <FormGroup>
              <FormLabel>Sensors</FormLabel>
              {data.water_level_sensors.map((sensor, idx) => (
                <FormControlLabel
                  key={idx}
                  control={
                    <Checkbox
                      color="primary"
                      checked={api.sensors.includes(sensor.id)}
                      onClick={(e) => handleSelectSensor(sensor.id)}
                    />
                  }
                  label={
                    <>
                      <Typography variant="body1">{sensor.id}</Typography>
                      <Chip
                        label={sensor.kind}
                        color={sensor.kind === "HIGH" ? "default" : "secondary"}
                        className={classes.chip}
                      />
                    </>
                  }
                />
              ))}
            </FormGroup>
          </FormControl>
        </Grid>
        <Grid item xs={12}>
          <FormControl className={classes.formControl}>
            <InputLabel>Pump</InputLabel>
            <Select
              labelId="input-pump-label"
              id="input-pump"
              value={api.pump}
              onChange={(e) => api.setPump(e.target.value)}
            >
              {data.pumps.map((s) => (
                <MenuItem key={s.id} value={s.id}>
                  {s.id}
                </MenuItem>
              ))}
            </Select>
          </FormControl>
          <FormControl className={classes.formControl}>
            <TextField
              id="input-rate"
              label="Fill Rate"
              helperText="Rate in mL/min"
              value={api.fillRate}
              onChange={(e) => api.setFillRate(e.target.value)}
            />
          </FormControl>
          <FormControl className={classes.formControl}>
            <TextField
              id="input-interval"
              label="Fill Interval"
              helperText="Rate in minutes"
              value={api.fillInterval}
              onChange={(e) => api.setFillInterval(e.target.value)}
            />
          </FormControl>
          <FormControl className={classes.formControl}>
            <TextField
              id="input-max-fill-volume"
              label="Max Fill Volume"
              helperText="Volume in mL"
              value={api.maxFillVolume}
              onChange={(e) => api.setMaxFillVolume(e.target.value)}
            />
          </FormControl>
        </Grid>
      </Grid>
    </CardContent>
  );
}

function SubmitStatus() {
  const api = useContext(Api);

  if (api.error) {
    console.log(api.error);
    return <Typography>Error submitting</Typography>;
  }
  return <></>;
}

function EditAutoTopOff() {
  const api = useContext(Api);

  return (
    <Card>
      <CardHeader title="Settings" />
      <Content />
      <CardActions>
        <Button color="primary" onClick={api.cancel}>
          Cancel
        </Button>
        <Button variant="contained" color="primary" onClick={api.submit}>
          Save
        </Button>
        <SubmitStatus />
      </CardActions>
    </Card>
  );
}

export default EditAutoTopOff;
