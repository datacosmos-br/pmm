import React from 'react';
import { Navigate, createBrowserRouter } from 'react-router-dom';
import { MainWithNav } from 'components/main/MainWithNav';
import Providers from 'Providers';
import { PMM_NEW_NAV_PATH } from 'lib/constants';
import { Redirect, SettingsRedirect } from 'components/redirect';

const router = createBrowserRouter(
  [
    {
      path: '',
      element: <Providers />,
      children: [
        {
          path: PMM_NEW_NAV_PATH,
          element: <MainWithNav />,
          children: [
            {
              path: '',
              element: <Navigate to="graph" />,
            },
            {
              path: 'updates',
              lazy: async () => ({
                Component: (await import('pages/updates')).Updates,
              }),
            },
            {
              path: 'updates/clients',
              lazy: async () => ({
                Component: (await import('pages/update-clients/UpdateClients'))
                  .UpdateClients,
              }),
            },
            {
              path: 'help',
              lazy: async () => ({
                Component: (await import('pages/help-center')).HelpCenter,
              }),
            },
            {
              path: 'settings/:tab?',
              lazy: async () => ({
                Component: (await import('pages/settings')).Settings,
              }),
            },
            {
              path: 'rta',
              children: [
                {
                  path: '',
                  lazy: async () => ({
                    Component: (await import('pages/rta/tab/RealtimeTab'))
                      .default,
                  }),
                },
                {
                  path: 'selection',
                  lazy: async () => ({
                    Component: (await import('pages/rta/selection'))
                      .RealtimeSelection,
                  }),
                },
                {
                  path: 'sessions',
                  lazy: async () => ({
                    Component: (await import('pages/rta/sessions'))
                      .RealtimeSessionsPage,
                  }),
                },
                {
                  path: 'overview',
                  lazy: async () => ({
                    Component: (
                      await import('pages/rta/overview/RealtimeOverview')
                    ).default,
                  }),
                },
              ],
            },
            // Fallback
            {
              path: 'graph/settings/:tab?',
              element: <SettingsRedirect />,
            },
            // Grafana routes are handled at the Main component level
            {
              path: 'graph/*',
              element: <React.Fragment />,
            },
            {
              path: '*',
              lazy: async () => ({
                Component: (await import('pages/not-found')).NotFoundPage,
              }),
            },
          ],
        },
        // Provide fallback for /next/* paths to redirect to the root path
        {
          path: '/next/*',
          element: <Redirect />,
        },
        {
          path: '*',
          element: <div>Not found!</div>,
        },
      ],
    },
  ],
  {
    basename: '/pmm-ui',
  }
);

export default router;
