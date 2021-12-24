Generating a list of calls from https://start.exactonline.nl/docs/HlpRestAPIResources.aspx?SourceAction=10:

``` js
var output = '';
var services = $('#ServiceFilter option').map(function() { return this.value.trim(); });
var calls = {};

$.each(services, function(i, service) {
    if (service == 'All') { return; }
    var trs = $('tr.' + service);
    output = output + "# " + service + "\n\n";

    trs.each(function() {
        var tr = $(this);
        var endpoint = tr.find('td').eq(1).text();
        var documentation = tr.find('td').eq(1).find('a').attr('href');
        documentation = 'https://start.exactonline.nl/docs/' + documentation;
        var resource = tr.find('td').eq(2).text();
        var methods = tr.find('td').eq(3).text().split(', ');

        output = output + "## " + endpoint + "\n\n";

        if (resource.startsWith('/api/') == false) {
            $.ajax(documentation, {async: false,  type: "POST"}).done(function(data) {
                var body = $(data);
                resource = body.find('#serviceUri').text();
            });
        }

        $.each(methods, function(j, method) {
            output = output + "- [ ] [" + method + " " + resource + "]" + "(" + documentation + ")" + "\n";
        });

        output = output + "\n";
    });
});
console.log(output);
```

Generate struct values from a call:

``` js
var types = [];
$('#referencetable>tbody>tr:nth-child(n+2):visible').each(function(i, tr) {
   var name = $.trim($(tr).children('td:nth-child(2)').text());
   var type = $.trim($(tr).children('td:nth-child(6)').text());
   var description = $.trim($(tr).children('td:nth-child(8), td:nth-child(9)').text());
   type = type.replace(/^Edm/, 'edm');
   type = type.replace(/^edm.Guid/, 'edm.GUID');

   types.push(name + ' ' + type + ' ' + '`json:"' + name + '"`' + ' // ' + description);
});
console.log(types.join("\n"));
```
