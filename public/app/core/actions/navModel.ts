import { NavModelItem } from '@grafana/ui';

export enum ActionTypes {
  UpdateNavIndex = 'UPDATE_NAV_INDEX',
}

export type Action = UpdateNavIndexAction;

export interface UpdateNavIndexAction {
  type: ActionTypes.UpdateNavIndex;
  payload: NavModelItem;
}

export const updateNavIndex = (item: NavModelItem): UpdateNavIndexAction => ({
  type: ActionTypes.UpdateNavIndex,
  payload: item,
});
