'use strict';

var webpack = require('webpack'),  
  HtmlWebpackPlugin = require('html-webpack-plugin'),
  path = require('path'),
  clientPath = path.join(__dirname, 'client');

module.exports = {  
  target: 'web',
  cache: true,
  entry: {
    module: path.join(clientPath, 'module.js'),
    common: ['react', 'react-router', 'alt']
  },
  resolve: {
    root: clientPath,
    extensions: ['', '.js'],
    modulesDirectories: ['node_modules', 'client']
  },
  output: {
    path: path.join(__dirname, 'public'),
    publicPath: '/static/',
    filename: '[name].js',
    library: ['Example', '[name]'],
    pathInfo: true
  },

  module: {
    loaders: [
      {test: /\.js?$/, exclude: /node_modules/, loader: 'babel?cacheDirectory'}
    ]
  },
  plugins: [
    new webpack.optimize.CommonsChunkPlugin('common', 'common.js'),
    new HtmlWebpackPlugin({
      inject: true,
      template: 'client/index.html'
    }),
    new webpack.NoErrorsPlugin()
  ],

  debug: true,
  devtool: 'eval-cheap-module-source-map',
  devServer: {
    contentBase: './tmp',
    historyApiFallback: true
  }
};
