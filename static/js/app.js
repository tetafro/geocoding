const krd = [45.04020, 38.97602];

function main() {
    map = L.map('map').setView(krd, 15);

    L.tileLayer('https://maps.wikimedia.org/osm-intl/{z}/{x}/{y}.png', {
        attribution: '<a href="https://wikimediafoundation.org/wiki/Maps_Terms_of_Use">Wikimedia</a>',
        minZoom: 1,
        maxZoom: 19
    }).addTo(map);

    const $input = $('#search');
    const $object = $('#object');

    $input.val('');

    $input
        .autocomplete({
            minLength: 3,
            source: function (request, response) {
                $.getJSON(
                    '/api/v1/places',
                    {name: request.term},
                    (resp) => {
                        return response(resp.data);
                    });
            },
            select: (event, ui) => {
                // Put marker on map and center it
                L.marker([ui.item.lat, ui.item.lon]).addTo(map);
                map.panTo([ui.item.lat, ui.item.lon]);
                $input.val(ui.item.display_name);

                // Display object
                $object.html('');
                let line;
                line = `<h4>${ui.item.display_name}</h4>`
                $('<div>').append(line).appendTo($object);
                for (field in ui.item) {
                    line = `<p><b>${field}:</b> ${ui.item[field]}</p>`;
                    $('<div>').append(line).appendTo($object);
                }

                event.preventDefault();
            }
        })
        .autocomplete('instance')._renderItem = function (ul, item) {
            return $('<li>')
                .append(`<div class="option">${item.display_name}</div>`)
                .appendTo(ul);
        };
}

// Run when page is loaded
document.addEventListener("DOMContentLoaded", function (event) {
    main();
});
