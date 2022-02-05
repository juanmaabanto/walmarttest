package command

// var (
// 	mongo_url = "mongodb+srv://root:A123a@develop.oh3sr.mongodb.net/test?retryWrites=true&w=majority"
// )

// func TestCreateProductHandler_Handle(t *testing.T) {
// 	description := uuid.New()
// 	type args struct {
// 		command CreateProduct
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    string
// 		wantErr bool
// 	}{
// 		{"Crear Nuevo Producto", args{CreateProduct{Brand: "marca test", Description: description.String(), Image: "http://image.com", Price: 10}}, "id", false},
// 		{"Error Producto existe", args{CreateProduct{Brand: "marca test", Description: description.String(), Image: "http://image.com", Price: 10}}, "", true},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			h := CreateProductHandler{
// 				repo: infrastructure.NewProductRepository(database.NewMongoConnection(context.Background(), "test", mongo_url), products.Product{}),
// 			}
// 			got, err := h.Handle(context.Background(), tt.args.command)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("CreateProductHandler.Handle() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if len(got) <= 0 && !tt.wantErr {
// 				t.Errorf("CreateProductHandler.Handle() = %v, id no generado", got)
// 			}
// 		})
// 	}
// }
