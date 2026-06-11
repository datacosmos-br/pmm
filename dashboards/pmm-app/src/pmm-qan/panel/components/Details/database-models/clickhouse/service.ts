import { apiRequest } from 'shared/components/helpers/api';

export default {
  getExplain(body, disableNotifications = false) {
    const requestBody = { clickhouse_explain: body };

    return apiRequest.post<any, any>('/v1/actions:startServiceAction', requestBody, disableNotifications);
  },

  getShowCreateTables: async () => null,
  getIndexes: async () => null,
};
