import { ChartOptions } from 'chart.js';
import {
  getThemePrimaryTextColor,
  getThemeSecondaryBackgroundColor,
} from 'shared/components/helpers/getPmmTheme';
import { GetDefaultOptionsProps } from './BarChart.types';

export const getDefaultOptions = ({
  options, orientation, barWidth, showLegend, theme,
}: GetDefaultOptionsProps): ChartOptions<'bar'> => {
  const indexAxis: 'x'|'y'| undefined = orientation && orientation === 'horizontal' ? 'y' : 'x';
  const primaryTextColor = getThemePrimaryTextColor(theme);
  const secondaryBackgroundColor = getThemeSecondaryBackgroundColor(theme);

  const defaultOptions = {
    indexAxis,
    responsive: true,
    maintainAspectRatio: false,
    barThickness: barWidth || 'flex',
    legend: {
      display: showLegend,
    },
    scales: {
      x: {
        grace: indexAxis === 'y' ? '1%' : 0,
      },
      y: {
        grace: indexAxis === 'x' ? '1%' : 0,
      },
    },
    animation: false as const,
    plugins: {
      datalabels: {
        color: primaryTextColor,
        anchor: 'end' as const,
        align: 'end' as const,
        formatter(_value, context) {
          if (context.dataset.dataInPersent) {
            const data = context.dataset.dataInPersent[context.dataIndex];

            return data ? `${data.toFixed(2)}%` : '';
          }

          return '';
        },
      },
      tooltip: {
        filter(tooltipItem) {
          return !!tooltipItem.formattedValue;
        },
        callbacks: {
          title() {
            return '';
          },
          label(context) {
            return context.formattedValue || '';
          },
          labelTextColor() {
            return primaryTextColor;
          },
        },
        backgroundColor: secondaryBackgroundColor,
        displayColors: false,
      },
      tooltips: {
        mode: 'index',
        intersect: false,
      },
    },
    resize: true,
    onResize(chart) {
      chart.update();
    },
  };

  return options ? {
    ...defaultOptions,
    ...options,
  } : defaultOptions;
};
