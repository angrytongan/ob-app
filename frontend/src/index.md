---
theme: dashboard
sql:
    csv: ./data/rand-xy.csv
---

# Observable app

This application demonstrates an Observable Framework application that houses a
Go backend and communicates to it using HTMX.

## Simple API calls

<div class="card">
    <h2>server time via htmx</h2>
    <button
        hx-get="/api/now"
        hx-target="#now"
        hx-swap="innerHTML">hit me</button>
    <span id="now"></span>
</div>

<!-- making sure deployment works with file attachments -->

```sql id=csv
SELECT r, x, y FROM csv
```

```js
const server = await fetch(`/api/random-points/${numPoints}`)
  .then((r) => {
    if (r.ok) {
      return r.json();
    }

    console.error(`received ${r.statusText}`)
    return [];
  })
  .catch((err) => {
    console.error(err);
    return []
  });
const data = [
  { points: csv, colour: 'green' },
  { points: server, colour: 'orange' },
];
```

```js
const dotPlot = (data) => resize((width) => {
  const config = {
    width,
    height: 200,
    x: {
      domain: [0, 100],
    },
    y: {
      domain: [0, 100],
    },
    marks: [
      Plot.axisX({
        ticks: d3.ticks(0, 100, 10),
      }),
      Plot.axisY({
        ticks: d3.ticks(0, 100, 10),
      }),
    ],
  };

  data.forEach((d) => {
    config.marks.push(
      Plot.dot(d.points, {
        x: "x",
        y: "y",
        r: "r",
        stroke: d.colour,
      })
    );
  });
    config.marks.push(
      Plot.dot([], {
        x: "x",
        y: "y",
        r: "r",
        stroke: 'blue'
      })
    );

  return Plot.plot(config)
});
```

<div class="card">

```js
const numPoints = view(Inputs.range([0, 100], { step: 10 }));
```

```js
display(dotPlot(data));
```

</div>
