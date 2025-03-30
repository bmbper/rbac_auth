import { defineConfig } from '@rsbuild/core';
import { pluginReact } from '@rsbuild/plugin-react';
import { pluginSvgr } from '@rsbuild/plugin-svgr';
import { pluginBasicSsl } from '@rsbuild/plugin-basic-ssl';
import { pluginSass } from '@rsbuild/plugin-sass';


export default {
  plugins: [pluginReact(), pluginSvgr(),pluginBasicSsl(),pluginSass()],
  html: {
    title: 'React Starter',
  },
  server: {
    port: 36000,

  },
  resolve: {
    alias: {
      '@src': './src',
    },
  },
};
