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
import { Avatar, LinearProgress, ListItemAvatar } from "@material-ui/core";

import VerticalAlignTopIcon from "@material-ui/icons/VerticalAlignTop";
import DeleteIcon from "@material-ui/icons/Delete";
import AddCircleIcon from "@material-ui/icons/AddCircle";

import IconAvatar from "./IconAvatar";

const DASHBOARD_QUERY = `query {
    auto_top_off {
        id
    }
}`;

const DELETE_MUTATION = `mutation DeleteAutoTopOff($id: ID!) {
    deleteAutoTopOff(id: $id)
}`;

function Item({ id, onDelete }) {
  return (
    <ListItem button component={RouterLink} to={`/ato/${id}`}>
      <ListItemAvatar>
        <Avatar
          alt={id}
          src={`https://avatars.dicebear.com/api/bottts/${id}.svg`}
        />
      </ListItemAvatar>
      <ListItemSecondaryAction>
        <Switch checked={true} />
        <IconButton edge="end" onClick={onDelete}>
          <DeleteIcon />
        </IconButton>
      </ListItemSecondaryAction>
    </ListItem>
  );
}

function Content() {
  const { loading, error, data, refetch } = useQuery(DASHBOARD_QUERY, {});
  const [deleteAutoTopOff] = useMutation(DELETE_MUTATION);

  if (loading)
    return (
      <CardContent>
        <LinearProgress />
      </CardContent>
    );
  if (error) return <CardContent>Failed to load ATO's</CardContent>;

  if (data.auto_top_off == null) return <></>;

  return (
    <CardContent>
      <List>
        {data.auto_top_off.map(({ id }) => (
          <Item
            id={id}
            key={id}
            onDelete={() => {
              deleteAutoTopOff({ variables: { id: id } });
              refetch();
            }}
          />
        ))}
      </List>
    </CardContent>
  );
}

function AutoTopOffs() {
  return (
    <Card>
      <CardHeader
        avatar={
          <IconAvatar>
            <VerticalAlignTopIcon />
          </IconAvatar>
        }
        title="Auto Top-Off"
      />
      <Content />
      <CardActions>
        <Link component={RouterLink} to="/ato/create">
          <Button startIcon={<AddCircleIcon />}>Add ATO</Button>
        </Link>
      </CardActions>
    </Card>
  );
}

export default AutoTopOffs;
