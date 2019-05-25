let cats = []
let catIndex = 0
let total = 0

$(function () {
    $('#prev').on('click', function () {
        if (catIndex > 0) {
            renderDishes(--catIndex)
        }
    })

    $('#next').on('click', function () {
        if (catIndex < cats.length - 1) {
            renderDishes(++catIndex)
        }
    })

    renderPage()
    renderCart()
})

function renderPage() {
    $.ajax({
        url: 'api/offers',
        type: 'get',
        success: function (res) {
            let offers = JSON.parse(res)
            let supplierId = offers[0].supplier_id
            renderSupplier(supplierId)
        }
    })
}

function renderSupplier(supplierId) {
    $.ajax({
        url: 'api/supplier/' + supplierId,
        type: 'get',
        success: function (res) {
            let supplier = JSON.parse(res)
            $('#supplier').text(supplier.name)
            $('#supplier-name').text(supplier.name)
            $('#supplier-addr').text(supplier.addr)
            $('#supplier-opened').text(supplier.opened)
            $('#supplier-phone').text(supplier.phone)

            renderCats(supplierId)
        }
    })
}

function renderCats(supplierId) {
    $.ajax({
        url: 'api/menu/' + supplierId,
        type: 'get',
        success: function (res) {
            cats = JSON.parse(res)

            for (let i = 0; i < cats.length; i++) {
                let html = '<a class="text-dark mx-2" href="">' + cats[i].name + '</a>'

                if (i < cats.length - 1) {
                    html += '/'
                }
                $('#cats').append(html)
            }

            renderDishes(catIndex)
        }
    })
}

function renderDishes(i) {
    $('#cat').text(cats[i].name)

    $('#bg-container').css('background-image', 'static/img/' + cats[i].img)

    $('#dishes').empty()

    for (let d of cats[i].dishes) {
        let html = '<div class="col-12 col-xl-6 px-5 py-3 add-item"><div class="row"><div class="col-12 col-lg-9"><span class="text-md text-very-strong">' + d.name + '</span><br /><span class="text-xs"><i>' + d.desc + '</i></span></div><div class="col-12 col-lg-3 text-right">+ <span class="text-white text-strong price" href="">' + Number.parseFloat(d.price).toFixed(2) + ' &euro;</span></div></div></div>'
        $('#dishes').append(html)
    }
}

function renderCart() {
    // Ajax request here
    let cart = {
        items: [
            {
                id: '7',
                name: 'Hamburger Royal K&auml;se',
                desc: 'Preis und Beschreibung folgen.',
                price: '3.9',
            },
            {
                id: '8',
                name: '9 Chicken McNuggets',
                desc: 'Preis und Beschreibung folgen.',
                price: '4.0',
            },
            {
                id: '9',
                name: '12 Chicken McNuggets',
                desc: 'Preis und Beschreibung folgen.',
                price: '3.2',
            }
        ],
        remark: 'Bitte mit Ketchup',
    }

    for (let i of cart.items) {
        let html = '<li class="list-inline text-xs text-strong" data-dish-id="' + i.id + '"><div class="row m-0"><div class="col-10 px-2 py-3">' + i.name + '</div><div class="col-2 py-3 remove-item text-center" data-dish-id="' + i.id + '">&#10005;</div></div></li>'
        $('#cart-items').append(html)
    }
    $('#cart-remark').val(cart.remark)
    $('#total').text(Number.parseFloat(total).toFixed(2))
}
