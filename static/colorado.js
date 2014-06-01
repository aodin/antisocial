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

var endings = ["th", "st", "nd", "rd"];

var hoverBar = "<h3><%- name %></h3><h1><%- rank %><sup><%- end %></h1><h4>of 78 neighborhoods</h4><table><tr><th>Population:</th><td><%- population %></td></tr><tr><th>Crimes:</th><td><%- crimes %></td></tr><tr><th>311 Calls:</th><td><%- calls %></td></tr><tr><th>Foreclosures:</th><td><%- foreclosures %></td></tr><tr><th>Liquor Licenses:</th><td><%- licenses %></td></tr><table>";

// Pass the google map to the view
var Map = Backbone.View.extend({
    el: '#map-canvas',
    initialize: function(options) {
        this.map = options.map;
        this.listenTo(this.collection, "reset", this.refresh);
        this.on('hover', this.hover, this);
        this.hoverbar = $('#hoverbar').hide();
    },
    refresh: function() {
      _.each(this.collection.models, function(m) {
        var h = new MapHood({model: m, map: this.map, t: this});
        h.render();
      }, this);
    },
    hover: function(m) {
      // TODO Just use a blacklist to include geometry in the attrs
      var r = m.get('rank');
      var n = Number(String(r).slice(-1));
      n = n > 3 ? 0 : n;
      var attrs = {
        name: m.get('name'),
        rank: m.get('rank'),
        end: endings[n],
        crimes: m.get('crimes'),
        calls: m.get('calls'),
        population: m.get('population'),
        licenses: m.get('licenses'),
        foreclosures: m.get('foreclosures'),
      };
      this.hoverbar.html(_.template(hoverBar, attrs)).show();
    }
});

var MapHood = Backbone.View.extend({
    initialize: function(options) {
        this.map = options.map;
        this.t = options.t;
    },
    render: function() {
        var rank = this.model.get('rank');
        var c = 'hsl(240, 0%, ' + (rank + 5) + '%)';
        var hoodOptions = {
          "strokeColor": "#4355B6",
          "strokeOpacity": 0.5,
          "strokeWeight": 2,
          "fillColor": c,
          "fillOpacity": 0.7
        }
        // convert the geom string to geo json
        var g = jQuery.parseJSON(this.model.get('geom'));
        var geo = new GeoJSON(g, hoodOptions);
        var parent = this.t;
        var model = this.model;
        google.maps.event.addListener(geo, 'mouseover', function() {
            // Trigger a new hover view
            parent.trigger('hover', model);
        });

        geo.setMap(this.map);
        return this;
    }
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
    zoom: 11,
    styles: styles,
    mapTypeControlOptions: {
        style: google.maps.MapTypeControlStyle.HORIZONTAL_BAR,
        position: google.maps.ControlPosition.RIGHT_BOTTOM
    },
    panControl: false,
    zoomControl: false,
    streetViewControl: false,
  };
  var map = new google.maps.Map(document.getElementById("map-canvas"),mapOptions);

  var hoods = new Hoods();
  var m = new Map({collection: hoods, map: map});
  hoods.fetch({reset: true});
});
