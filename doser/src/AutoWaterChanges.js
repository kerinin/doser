import React from "react";

import Card from "@material-ui/core/Card";
import CardContent from "@material-ui/core/CardContent";
import CardActions from "@material-ui/core/CardActions";
import CardHeader from "@material-ui/core/CardHeader";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import ListItemSecondaryAction from "@material-ui/core/ListItemSecondaryAction";
import Button from "@material-ui/core/Button";
import Switch from "@material-ui/core/Switch";
import IconButton from "@material-ui/core/IconButton";
import { Link as RouterLink } from "react-router-dom";
import Link from "@material-ui/core/Link";
import { useQuery } from "graphql-hooks";
import { useMutation } from "graphql-hooks";
import { Avatar, ListItemAvatar } from "@material-ui/core";

import DeleteIcon from "@material-ui/icons/Delete";
import AddCircleIcon from "@material-ui/icons/AddCircle";
import AutorenewIcon from "@material-ui/icons/Autorenew";

import IconAvatar from "./IconAvatar";

const DASHBOARD_QUERY = `query {
    auto_water_changes {
        id
        enabled
    }
}`;

const SET_ENABLED = `mutation SetAWCEnabled($id: ID!, $enabled: Boolean!) {
  setAutoWaterChangeEnabled(id: $id, enabled: $enabled)
}`;

const DELETE_MUTATION = `mutation DeleteAutoWaterChange($id: ID!) {
    deleteAutoWaterChange(id: $id)
}`;

function Item({ id, enabled, onDelete, refresh }) {
  const [setEnabled] = useMutation(SET_ENABLED);

  function toggleEnabled() {
    setEnabled({ variables: { id: id, enabled: !enabled } }).then(() => {
      console.log("refreshing");
      refresh();
    });
  }

  return (
    <ListItem button component={RouterLink} to={`/awc/${id}`}>
      <ListItemAvatar>
        <Avatar
          alt={id}
          src={`https://avatars.dicebear.com/api/bottts/${id}.svg`}
        />
      </ListItemAvatar>
      <ListItemSecondaryAction>
        <Switch checked={enabled} onClick={() => toggleEnabled()} />
        <IconButton edge="end" onClick={onDelete}>
          <DeleteIcon />
        </IconButton>
      </ListItemSecondaryAction>
    </ListItem>
  );
}

function Content() {
  const { loading, error, data, refetch } = useQuery(DASHBOARD_QUERY, {});
  const [deleteAutoWaterChange] = useMutation(DELETE_MUTATION);

  if (loading) return <CardContent />;
  if (error) return <CardContent>Failed to load AWC's</CardContent>;

  if (data.auto_water_changes == null) return <></>;

  return (
    <CardContent>
      <List>
        {data.auto_water_changes.map(({ id, enabled }) => (
          <Item
            id={id}
            key={id}
            enabled={enabled}
            onDelete={() => {
              deleteAutoWaterChange({ variables: { id: id } });
              refetch();
            }}
            refresh={refetch}
          />
        ))}
      </List>
    </CardContent>
  );
}

function AutoWaterChanges() {
  return (
    <Card>
      <CardHeader
        avatar={
          <IconAvatar>
            <AutorenewIcon />
          </IconAvatar>
        }
        title="Auto Water Change"
      />
      <Content />
      <CardActions>
        <Link component={RouterLink} to="/awc/create">
          <Button startIcon={<AddCircleIcon />}>Add AWC</Button>
        </Link>
      </CardActions>
    </Card>
  );
}

export default AutoWaterChanges;
