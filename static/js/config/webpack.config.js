const webpack = require( 'webpack' );
const UglifyJsPlugin = require( 'uglifyjs-webpack-plugin' );
const paths = require( './paths' );

module.exports = {
  devtool: 'source-map',
  entry: paths.appIndexJs,
  module: {
    rules: [
      {
        test: /\.(js|jsx)$/,
        exclude: /node_modules/,
        use: ['babel-loader']
      },
      {
        test: /\.js$/,
        exclude: /node_modules/,
        use: ['babel-loader', 'eslint-loader']
      }
    ]
  },
  resolve: {
    extensions: [
      '*', '.js', '.jsx'
    ]
  },
  output: {
    path: paths.appDist,
    publicPath: '/',
    filename: './cryptopapers-graphql-api.min.js',
    sourceMapFilename: './cryptopapers-graphql-api.min.js.map'
  },
  optimization: {
    minimizer: [
      new UglifyJsPlugin( {
        sourceMap: true
      } )
    ]
  },
};