import React, { Suspense, lazy } from 'react';
import { PanelPlugin, PanelProps } from '@grafana/data';

const QueryAnalyticsPanel = lazy(() => import('pmm-qan/panel/QueryAnalytics'));

const QueryAnalyticsLazyPanel = (props: PanelProps) => React.createElement(
  Suspense,
  { fallback: null },
  React.createElement(QueryAnalyticsPanel, props),
);

export const plugin = new PanelPlugin(QueryAnalyticsLazyPanel);
