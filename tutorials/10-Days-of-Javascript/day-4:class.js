class Polygon {
    constructor (obj)
    {
        var polygon_perimeter = 0;
        for (var index in obj)
        {
            polygon_perimeter += obj[index];
        }

        this.my_perimeter = polygon_perimeter;
    }

    perimeter()
    {
        return this.my_perimeter;
    }
}

const rectangle = new Polygon([10, 20, 10, 20]);
const square = new Polygon([10, 10, 10, 10]);
const pentagon = new Polygon([10, 20, 30, 40, 43]);

console.log(rectangle.perimeter());
console.log(square.perimeter());
console.log(pentagon.perimeter());