import { GrafanaTheme, GrafanaTheme2 } from '@grafana/data';

interface TableTheme {
  backgroundColor: string;
  borderColor: string;
  headerBackground: string;
  textColor: string;
  selectedRowColor: string;
  sortIconColor: string;
}

interface DropdownTheme {
  bg: string;
  text: string;
  hoverBg: string;
  bgSmallText: string;
}

interface Themes {
  mainTextColor: string;
  table: TableTheme;
  dropdown: DropdownTheme;
}

/**
 * Returns PMM-specific theme values based on Grafana theme.
 * Provides colors for tables, dropdowns, and text.
 * Defaults to light theme if Grafana theme is not initialized.
 *
 * @param theme - The current Grafana theme
 * @returns Theme configuration object with colors for different UI elements
 */
type PmmGrafanaTheme = GrafanaTheme | GrafanaTheme2 | undefined;

type UnknownRecord = Record<string, unknown>;

export interface ThemePrimaryColor {
  main: string;
  shade: string;
  contrastText: string;
}

const isRecord = (value: unknown): value is UnknownRecord => typeof value === 'object' && value !== null;

const getRequiredThemeValue = (theme: unknown, path: readonly string[], description: string): unknown => {
  let value = theme;

  path.forEach((key) => {
    if (!isRecord(value)) {
      throw new Error(`${description} is required`);
    }

    value = value[key];
  });

  return value;
};

const getRequiredThemeString = (theme: unknown, path: readonly string[], description: string): string => {
  const value = getRequiredThemeValue(theme, path, description);

  if (typeof value !== 'string') {
    throw new Error(`${description} is required`);
  }

  return value;
};

export const getThemePrimaryBackgroundColor = (theme: unknown): string => (
  getRequiredThemeString(theme, ['colors', 'background', 'primary'], 'Grafana theme colors.background.primary')
);

export const getThemeSecondaryBackgroundColor = (theme: unknown): string => (
  getRequiredThemeString(theme, ['colors', 'background', 'secondary'], 'Grafana theme colors.background.secondary')
);

export const getThemeCanvasBackgroundColor = (theme: unknown): string => (
  getRequiredThemeString(theme, ['colors', 'background', 'canvas'], 'Grafana theme colors.background.canvas')
);

export const getThemeActionFocusColor = (theme: unknown): string => (
  getRequiredThemeString(theme, ['colors', 'action', 'focus'], 'Grafana theme colors.action.focus')
);

export const getThemePrimaryTextColor = (theme: unknown): string => (
  getRequiredThemeString(theme, ['colors', 'text', 'primary'], 'Grafana theme colors.text.primary')
);

export const getThemePrimaryColor = (theme: unknown): ThemePrimaryColor => ({
  main: getRequiredThemeString(theme, ['colors', 'primary', 'main'], 'Grafana theme colors.primary.main'),
  shade: getRequiredThemeString(theme, ['colors', 'primary', 'shade'], 'Grafana theme colors.primary.shade'),
  contrastText: getRequiredThemeString(
    theme,
    ['colors', 'primary', 'contrastText'],
    'Grafana theme colors.primary.contrastText',
  ),
});

const getLightMainTextColor = (theme: PmmGrafanaTheme): string => {
  if (!theme) {
    return '#000000';
  }

  const text = getRequiredThemeValue(theme, ['colors', 'text'], 'Grafana theme colors.text');

  if (typeof text === 'string') {
    return text;
  }

  return getRequiredThemeString(text, ['primary'], 'Grafana theme colors.text.primary');
};

const getMainTextColor = (theme: PmmGrafanaTheme, isLight: boolean): string => {
  if (!isLight) {
    return 'rgba(255, 255, 255, 0.8)';
  }

  if (!theme) {
    return '#000000';
  }

  return getLightMainTextColor(theme);
};

export const getPmmTheme = (theme: PmmGrafanaTheme): Themes => {
  const isLight = theme?.isLight ?? true;
  const mainTextColor = getMainTextColor(theme, isLight);

  const backgroundColor = isLight ? '#f7f8fa' : '#0b0c0e';
  const borderColor = isLight ? mainTextColor : '#292929';
  const headerBackground = isLight ? '#dedfe1' : '#202226';
  const selectedRowColor = isLight ? 'deepskyblue' : '#234682';
  const sortIconColor = 'deepskyblue';
  const bg = isLight ? '#ffffff' : '#262626';
  const bgSmallText = isLight ? 'transparent' : '#646464';
  const text = isLight ? '#000000' : '#d8d9da';
  const hoverBg = isLight ? '#f5f5f5' : '#333333';

  return {
    mainTextColor,
    table: {
      backgroundColor,
      borderColor,
      headerBackground,
      textColor: mainTextColor,
      selectedRowColor,
      sortIconColor,
    },
    dropdown: {
      bg,
      text,
      hoverBg,
      bgSmallText,
    },
  };
};

/**
 * Applies PMM theme CSS variables to document.body.
 * These variables are used throughout QAN for consistent theming of dropdowns,
 * backgrounds, and text colors in both light and dark modes.
 *
 * Variables set:
 * - --qan-dropdown-bg: Background color for AntD Select dropdowns
 * - --qan-dropdown-text: Text color for dropdown options
 * - --qan-dropdown-hover-bg: Background color for hovered/selected dropdown options
 * - --page-background: Main background color for QAN panel
 * - --main-text-color: Primary text color
 *
 * @param grafanaTheme - The current Grafana theme
 */
export const applyPmmCssVariables = (grafanaTheme: PmmGrafanaTheme): void => {
  const pmmTheme = getPmmTheme(grafanaTheme);

  if (typeof document === 'undefined') {
    return;
  }

  const root = document.body;

  // QAN dropdown variables
  root.style.setProperty('--qan-dropdown-bg', pmmTheme.dropdown.bg);
  root.style.setProperty('--qan-dropdown-text', pmmTheme.dropdown.text);
  root.style.setProperty('--qan-dropdown-hover-bg', pmmTheme.dropdown.hoverBg);
  root.style.setProperty('--qan-dropdown-bgSmallText', pmmTheme.dropdown.bgSmallText);

  // Page background and text color
  root.style.setProperty('--page-background', pmmTheme.table.backgroundColor);
  root.style.setProperty('--main-text-color', pmmTheme.mainTextColor);
  root.style.setProperty('--main-qan-font', 'inherit');
};
