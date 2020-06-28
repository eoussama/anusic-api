const path = require('path');

module.exports = {
  entry: './src/index.mjs',
  mode: 'production',
  module: {
    rules: [
      {
        test: /\.m?js$/,
        exclude: /(node_modules|bower_components)/,
        use: {
          loader: 'babel-loader',
          options: {
            presets: ['@babel/preset-env']
          }
        }
      }
    ],
  },
  output: {
    filename: 'index.min.js',
    path: path.resolve(__dirname, 'dist'),
  },
};
