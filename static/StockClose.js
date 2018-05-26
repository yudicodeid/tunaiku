var app = this.app = this.app || {};

(function(){

    var StockClose = function() {

        this.data = [];
        this.container = '.stock-close-list';
        this.btnAdd = '.btn-add-stock-close';
        this.addContainer = '#stockCloseModal';
        this.btnDelete = '.btn-delete-stock';
        this.onListUpdated = null;

        var that = this;

        this.assignDeleteEvent = function() {

            $(this.btnDelete).on('click', function(){
                var id = $(this).attr('data-id');
                that.delete_(id);

            });

        }


        $(this.btnAdd).click(function(){
            that.add();
        });

        this.setOnListUpdated = function(handler) {
            this.onListUpdated = handler;
        }

        this.list = function() {

            var that = this;

            $.ajax({
                url : '/list',
                method : 'get',
                dataType :'json',
                success : function(d) {

                    console.log(d);
                    if(d.Status == true) {

                        that.data = d.Entities;
                        //that.onListUpdated(that.data);

                        that.render();
                    }
                    else {
                        console.log(d.Messsage)
                    }

                }
            });
        }



        this.add = function() {

            var that = this;
            $.ajax({
                url : '/',
                method : 'POST',
                data :  {
                    StockDate : $('#StockDate').val(),
                    Open : $('#Open').val(),
                    High : $('#High').val(),
                    Low : $('#Low').val(),
                    Close : $('#Close').val(),
                    VolumeTrade : $('#VolumeTrade').val()
                },
                contentType: 'application/x-www-form-urlencoded',
                dataType :'json',
                success : function(d) {
                    if(d.Status === true) {

                        $('#StockDate').val('');
                        $('#Open').val(0);
                        $('#High').val(0);
                        $('#Low').val(0);
                        $('#Close').val(0);
                        $('#VolumeTrade').val(0);

                        $(that.addContainer).modal('hide');
                        that.list();
                    }
                    else {
                        alert(d.Message);
                    }
                }

            });

        }


        this.delete_ = function(id) {

            var that = this;
            $.ajax({
                url : '/',
                method : 'PATCH',
                data :  {
                    ID : id
                },
                contentType: 'application/x-www-form-urlencoded',
                dataType :'json',
                success : function(d) {
                    if(d.Status === true) {
                        that.list();
                    }
                    else {
                        alert(d.Message);
                    }
                }
            });

        }




        this.render = function() {

            var str = '';
            $(this.container).empty();

            if(this.data == null) return;

            for(var i=0; i< this.data.length; i++) {
                var d = this.data[i];
                str +='<tr data-id="' + d['ID'] + '">' +
                    '<td><button class="btn btn-danger btn-delete-stock" data-id="' + d['ID'] + '">Delete</button></td>' +
                    '<td>' + d['StockDate'] + '</td>' +
                    '<td>' + d['Open'] + '</td>' +
                    '<td>' + d['High'] + '</td>' +
                    '<td>' + d['Low'] + '</td>' +
                    '<td>' + d['Close'] + '</td>' +
                    '<td>' + d['VolumeTrade'] + '</td>' +
                    '<td>' + d['Action'] + ( d['Max'] == true ? '(Max)' :'') + '</td>' +
                    '</tr>';
            }

            $(this.container).append(str);

            this.assignDeleteEvent();

        }

    }

    app.StockClose = StockClose;

})();

