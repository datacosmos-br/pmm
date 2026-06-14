import React, { HTMLProps, useCallback } from 'react';
import { GrafanaTheme2 } from '@grafana/data';
import { useStyles2 } from '@grafana/ui';
import { css, cx } from '@emotion/css';
import {
  getThemeActionFocusColor,
  getThemeCanvasBackgroundColor,
  getThemePrimaryColor,
  getThemePrimaryTextColor,
} from 'shared/components/helpers/getPmmTheme';

export interface CheckboxProps extends Omit<HTMLProps<HTMLInputElement>, 'value'> {
  label?: string;
  value?: boolean;
}

export const getFocusCss = (theme: GrafanaTheme2) => css`
  outline: 2px dotted transparent;
  outline-offset: 2px;
  box-shadow: 0 0 0 2px ${getThemeCanvasBackgroundColor(theme)}, 0 0 0px 4px ${getThemeActionFocusColor(theme)};
  transition: all 0.2s cubic-bezier(0.19, 1, 0.22, 1);
`;

export const getLabelStyles = (theme: GrafanaTheme2) => ({
  label: css`
      font-size: ${theme.typography.bodySmall.fontSize};
      font-weight: ${theme.typography.fontWeightMedium};
      line-height: ${theme.typography.bodySmall.lineHeight};
      margin: 0 0 ${theme.spacing(0.5)} 0;
      padding: 0;
      color: ${getThemePrimaryTextColor(theme)};
      max-width: 480px;
    `,
});

export const getCheckboxStyles = (theme: GrafanaTheme2) => {
  const labelStyles = getLabelStyles(theme);
  const checkboxSize = '16px';
  const { main: primaryMain, shade: primaryShade, contrastText: primaryContrastText } = getThemePrimaryColor(theme);

  return {
    label: cx(
      labelStyles.label,
      css`
        padding-left: ${theme.spacing(1)};
      `,
    ),
    wrapper: css`
      position: relative;
      padding-left: ${checkboxSize};
      vertical-align: middle;
    `,
    input: css`
      position: absolute;
      top: 0;
      left: 0;
      width: 100%;
      height: 100%;
      opacity: 0;
      &:focus + span {
        ${getFocusCss(theme)}
      }
      &:checked + span {
        background: ${primaryMain};
        border: none;
        &:hover {
          background: ${primaryShade};
        }
        &:after {
          content: '';
          position: absolute;
          left: 5px;
          top: 1px;
          width: 6px;
          height: 12px;
          border: solid ${primaryContrastText};
          border-width: 0 3px 3px 0;
          transform: rotate(45deg);
        }
      }
    `,
    checkmark: css`
      display: inline-block;
      width: ${checkboxSize};
      height: ${checkboxSize};
      border-radius: ${theme.shape.radius.sm};
      margin-right: ${theme.spacing(1)};
      background: ${theme.components.input.background};
      border: 1px solid ${theme.components.input.borderColor};
      position: absolute;
      top: 1px;
      left: 0;
      &:hover {
        cursor: pointer;
        border-color: ${theme.components.input.borderHover};
      }
    `,
  };
};

export const Checkbox = React.forwardRef<HTMLInputElement, CheckboxProps>(
  ({
    label, value, onChange, disabled, ...inputProps
  }, ref) => {
    const handleOnChange = useCallback(
      (e: React.ChangeEvent<HTMLInputElement>) => {
        if (onChange) {
          onChange(e);
        }
      },
      [onChange],
    );
    const styles = useStyles2(getCheckboxStyles);

    return (
      <label className={styles.wrapper}>
        <input
          type="checkbox"
          className={styles.input}
          checked={value}
          disabled={disabled}
          onChange={handleOnChange}
          {...inputProps}
          ref={ref}
        />
        <span className={styles.checkmark} />
        {label && <span className={styles.label}>{label}</span>}
      </label>
    );
  },
);

Checkbox.displayName = 'Checkbox';
