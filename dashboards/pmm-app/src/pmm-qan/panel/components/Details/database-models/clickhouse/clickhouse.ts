import ClickHouseDatabaseService from './service';

export const clickhouseMethods = {
  getShowCreateTables: async () => null,
  getIndexes: async () => null,
  getExplainJSON: async () => null,
  getExplain: async ({ example }, disableNotifications = false) => {
    const payload = {
      service_id: example.service_id,
      query: example.example,
      explain_type: '',
    };

    const result = await ClickHouseDatabaseService.getExplain(payload, disableNotifications);

    return result.clickhouse_explain.action_id;
  },
};
