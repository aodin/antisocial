var colors = [
  "#f0585e", // Red
  "#5899d2", // Blue
  "#78c269", // Green
  "#f9a65a", // Orange
  "#9d65aa", // Purple
  "#4fc99d", // Teal
  "#cc6f57", // Magenta
  "#d67eb2", // Lavender
  "#5c68aa", // Blue 2
  "#5ce160", // Toxic Green
  "#d6d67e", // Ugly yellow
];

var Hood = Backbone.Model.extend({
    urlRoot: '/api',
});

var Hoods = Backbone.Collection.extend({
    model: Hood,
    url: '/api',
});

// Pass the google map to the view
var Map = Backbone.View.extend({
    el: '#map-canvas',
    initialize: function(options) {
        console.log('init:', options)
        this.map = options.map;
        this.listenTo(this.collection, "reset", this.refresh);
    },
    refresh: function() {
        _.each(this.collection.models, function(m, i) {

            var rank = m.get('rank');

            // var c = 'hsl(240, 0%, ' + (78 - rank) + '%)';
            var c = 'hsl(240, 0%, ' + (rank + 5) + '%)';
            console.log('c', c);

            var hoodOptions = {
              // "strokeColor": colors[i % colors.length],
              "strokeColor": "#4355B6",
              "strokeOpacity": 0.5,
              "strokeWeight": 2,
              "fillColor": c,
              // "fillColor": colors[i % colors.length],
              "fillOpacity": 0.7
            }
            // convert the geom string to geo json
            var g = jQuery.parseJSON(m.get('geom'));
            console.log('g:', g);

            var geo = new GeoJSON(g, hoodOptions);
            console.log(geo, geo.error);
            geo.setMap(this.map);
        }, this);
    },
});


$(function() {

    var styles = [
    {
      "featureType": "water",
      "stylers": [
        { "color": "#808082" },
        { "hue": "#003bff" },
        { "saturation": 46 },
        { "lightness": 30 }
      ]
    },{
      "featureType": "road",
      "stylers": [
        { "saturation": -100 },
        { "lightness": 60 }
      ]
    },{
      "featureType": "poi",
      "stylers": [
        { "hue": "#00ff66" }
      ]
    },{
    }
  ];

  var mapOptions = {
    // center: new google.maps.LatLng(39.739167, -104.984722),
    center: new google.maps.LatLng(39.719167, -104.944722),
    zoom: 12,
    styles: styles
  };
  var map = new google.maps.Map(document.getElementById("map-canvas"),mapOptions);

  var hoods = new Hoods();
  var m = new Map({collection: hoods, map: map});

  // Create the bounds
  // var sw = new google.maps.LatLng(39.74157597976737, -104.97832761370853);
  // var ne = new google.maps.LatLng(39.72325901999735, -104.93176446521);
  // var bounds = new google.maps.LatLngBounds(sw, ne);

    hoods.fetch({reset: true});
});



// if (congress.error) {
//   // Handle the error.
//   console.log('ERROR:', congress.error);
// } else {
//   congress.setMap(map);
// }

// Focus on bounds
// map.setCenter(origin);
// map.fitBounds(bounds);


// Load the zones
// for (var i = 0, len = zones.length; i < len; i++) {

// var zoneOptions = {
//   "strokeColor": colors[i],
//   "strokeOpacity": 0.75,
//   "strokeWeight": 1,
//   "fillColor": colors[i],
//   "fillOpacity": 0.5
// }

// var zone = new GeoJSON(zones[i], zoneOptions);
// if (zone.error) {
// // Handle the error.
//   console.log('ERROR:', zone.error);
// } else {
//   zone.setMap(map);
// }
// zone.setMap(map);
// }

// google.maps.event.addListener(map, 'center_changed', function() {
// console.log(map.getBounds());
// });
