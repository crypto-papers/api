package graphiql

var graphiqlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>Graphiql</title>
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/{{ .GraphiqlVersion }}/graphiql.css" />
	<script src="//cdn.jsdelivr.net/npm/es6-promise@{{ .Es6PromiseVersion }}/dist/es6-promise.auto.min.js"></script>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/{{ .FetchVersion }}/fetch.min.js"></script>
  <script src="https://cdnjs.cloudflare.com/ajax/libs/react/{{ .ReactVersion }}/umd/react.production.min.js"></script>
  <script src="https://cdnjs.cloudflare.com/ajax/libs/react-dom/{{ .ReactVersion }}/umd/react-dom.production.min.js"></script>
  <script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/{{ .GraphiqlVersion }}/graphiql.js"></script>
</head>
<body>
<div id="graphiql" style="height: 100vh;">Loading...</div>
<script>
	function graphQLFetcher(graphQLParams) {
		return fetch(window.location.origin + '/graphql', {
			method: 'post',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(graphQLParams),
		}).then(response => response.json());
	}

	ReactDOM.render(
		React.createElement(GraphiQL, {fetcher: graphQLFetcher}),
		document.getElementById("graphiql")
	);
</script>
</body>
</html>
`
