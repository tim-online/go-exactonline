``` javascript
var types = [];
$('#referencetable>tbody>tr:nth-child(n+2)').each(function(i, tr) {
   var name = $.trim($(tr).children('td:nth-child(2)').text());
   var type = $.trim($(tr).children('td:nth-child(5)').text());
   type = type.replace(/^Edm/, 'edm');
   type = type.replace(/^edm.Guid/, 'edm.GUID');

   types.push(name + ' ' + type + ' ' + '`json:"' + name + '"`');
});
console.log(types.join("\n"));
```
