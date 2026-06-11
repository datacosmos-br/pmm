import { css } from '@emotion/css';
import { GrafanaTheme } from '@grafana/data';

export const getStyles = (theme: GrafanaTheme) => ({
  errorMessage: css`
    color: ${theme.palette.redBase};
    font-size: ${theme.typography.size.sm};
    height: ${theme.typography.size.sm};
    line-height: ${theme.typography.lineHeight.sm};
    margin-top: ${theme.spacing.sm};
    margin-bottom: ${theme.spacing.xs};
  `,
  input: css`
    input {
      /* TODO: remove once using only platform-core components */
      min-height: 37px;
    }
  `,
});
