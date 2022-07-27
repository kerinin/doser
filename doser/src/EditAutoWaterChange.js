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
    pumps {
        id
        name
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

  if (loading) return <CardContent />;
  if (error) return <CardContent>Failed to load pumps</CardContent>;

  if (data.pumps == null)
    return <CardContent>No pumps defined - create one first</CardContent>;

  return (
    <CardContent>
      <Grid container spacing={2} className={classes.grid}>
        <Grid item xs={12}>
          <FormControl className={classes.formControl}>
            <InputLabel>Fresh Pump</InputLabel>
            <Select
              labelId="input-fresh-pump-label"
              id="input-fresh-pump"
              value={api.freshPump}
              onChange={(e) => api.setFreshPump(e.target.value)}
            >
              {data.pumps.map((s) => (
                <MenuItem key={s.id} value={s.id}>
                  {s.name || s.id}
                </MenuItem>
              ))}
            </Select>
          </FormControl>
          <FormControl className={classes.formControl}>
            <InputLabel>Waste Pump</InputLabel>
            <Select
              labelId="input-waste-pump-label"
              id="input-waste-pump"
              value={api.wastePump}
              onChange={(e) => api.setWastePump(e.target.value)}
            >
              {data.pumps.map((s) => (
                <MenuItem key={s.id} value={s.id}>
                  {s.name || s.id}
                </MenuItem>
              ))}
            </Select>
          </FormControl>
          <FormControl className={classes.formControl}>
            <TextField
              id="input-exchange-rate"
              label="Exchange Rate"
              helperText="L/day"
              value={api.exchangeRate}
              onChange={(e) => api.setExchangeRate(e.target.value)}
            />
          </FormControl>
          <FormControl className={classes.formControl}>
            <TextField
              id="input-salinity-adjustment"
              label="Salinity Adjustment"
              helperText="mL/day not removed from tank. Salinity increases with larger values."
              value={api.salinityAdjustment}
              onChange={(e) => api.setSalinityAdjustment(e.target.value)}
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

function EditAutoWaterChange() {
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

export default EditAutoWaterChange;
