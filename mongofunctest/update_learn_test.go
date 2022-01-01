package mongofunctest

// func TestUpdateHero_MultipleSet(t *testing.T) {
// 	ctx := context.Background()

// 	filter := mongoquery.Init(
// 		mongoquery.Equal("damage", ROUND),
// 	)
// 	update := mongoquery.Init(
// 		mongoquery.Set(mongoquery.PairSetter{
// 			FieldName: "damage",
// 			Value:     -ROUND,
// 		}),
// 		mongoquery.Set(mongoquery.PairSetter{
// 			FieldName: "name",
// 			Value:     "name",
// 		}),
// 	)

// 	if err := mongorely.UpdateMany(ctx, COLLECTION_NAME, filter, update); err != nil {
// 		t.Error("mongoquery Set failed 1", err)
// 		return
// 	}

// 	filter2 := mongoquery.Init(
// 		mongoquery.Equal("damage", -ROUND),
// 	)
// 	if count, err := mongorely.Count(ctx, COLLECTION_NAME, filter2); err != nil {
// 		t.Error("mongoquery Set failed 2", err)
// 		return
// 	} else if count != int64(ROUND) {
// 		t.Error("mongoquery Set wrong 3", count, "!=", ROUND)
// 		return
// 	}

// 	update2 := mongoquery.Init(
// 		mongoquery.Set(mongoquery.PairSetter{
// 			FieldName: "damage",
// 			Value:     ROUND,
// 		}),
// 	)

// 	if err := mongorely.UpdateMany(ctx, COLLECTION_NAME, filter2, update2); err != nil {
// 		t.Error("mongoquery Set failed 4", err)
// 		return
// 	}

// 	if count, err := mongorely.Count(ctx, COLLECTION_NAME, filter); err != nil {
// 		t.Error("mongoquery Set failed 5", err)
// 		return
// 	} else if count != int64(ROUND) {
// 		t.Error("mongoquery Set wrong 6", count, "!=", ROUND)
// 		return
// 	}
// }

// func TestUpdateHero_NestedSet(t *testing.T) {
// 	ctx := context.Background()

// 	filter := mongoquery.Init(
// 		mongoquery.Equal("damage", ROUND),
// 	)
// 	update := bson.M{
// 		"$set": bson.M{
// 			"subhero": bson.M{
// 				"1": bson.M{
// 					"amount": 1,
// 				},
// 			},
// 		},
// 	}

// 	if err := mongorely.UpdateMany(ctx, COLLECTION_NAME, filter, update); err != nil {
// 		t.Error("mongoquery Set failed 1", err)
// 		return
// 	}

// 	update2 := bson.M{
// 		"$set": bson.M{
// 			"subhero": bson.M{
// 				"2": bson.M{
// 					"amount": 1,
// 				},
// 			},
// 		},
// 	}

// 	if err := mongorely.UpdateMany(ctx, COLLECTION_NAME, filter, update2); err != nil {
// 		t.Error("mongoquery Set failed 1", err)
// 		return
// 	}
// }

// func BenchmarkUpdateHero_MultipleSet1(b *testing.B) {
// 	b.Run("test1", func(b *testing.B) {
// 		for n := 0; n < b.N; n++ {
// 			ctx := context.Background()

// 			filter := mongoquery.Init(
// 				mongoquery.Equal("damage", ROUND),
// 			)
// 			update := mongoquery.Init(
// 				mongoquery.Set(mongoquery.PairSetter{
// 					FieldName: "omit",
// 					Value:     1,
// 				}),
// 				mongoquery.Set(mongoquery.PairSetter{
// 					FieldName: "name",
// 					Value:     "name",
// 				}),
// 			)

// 			if err := mongorely.UpdateMany(ctx, COLLECTION_NAME, filter, update); err != nil {
// 				b.Error("mongoquery Set failed 1", err)
// 				return
// 			}
// 		}
// 	})

// 	b.Run("test3", func(b *testing.B) {
// 		for n := 0; n < b.N; n++ {
// 			ctx := context.Background()

// 			filter := mongoquery.Init(
// 				mongoquery.Equal("damage", ROUND),
// 			)
// 			update := mongoquery.Init(
// 				mongoquery.Set(mongoquery.PairSetter{
// 					FieldName: "omit",
// 					Value:     1,
// 				}),
// 				mongoquery.Set(mongoquery.PairSetter{
// 					FieldName: "name",
// 					Value:     "name",
// 				}),
// 			)

// 			if err := mongorely.UpdateMany(ctx, COLLECTION_NAME, filter, update); err != nil {
// 				b.Error("mongoquery Set failed 1", err)
// 				return
// 			}
// 		}
// 	})

// 	b.Run("test2", func(b *testing.B) {
// 		for n := 0; n < b.N; n++ {
// 			ctx := context.Background()

// 			filter := mongoquery.Init(
// 				mongoquery.Equal("damage", ROUND),
// 			)
// 			update := mongoquery.Init(
// 				mongoquery.Set(mongoquery.PairSetter{
// 					FieldName: "omit",
// 					Value:     1,
// 				}, mongoquery.PairSetter{
// 					FieldName: "name",
// 					Value:     "name",
// 				}),
// 			)

// 			if err := mongorely.UpdateMany(ctx, COLLECTION_NAME, filter, update); err != nil {
// 				b.Error("mongoquery Set failed 1", err)
// 				return
// 			}
// 		}
// 	})
// }
