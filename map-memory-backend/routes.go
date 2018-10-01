/* Copyright Fengchang Xie & Huifeng Zhang, 2018 All Rights Reserved */
package main

/**
 * Created by Fengchang on 2018/5/30 12:53.
 */

func initializeRoutes() {
	// Use the setUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	// router.Use(setUserStatus())
	// Handle the index route

	router.GET("/api/v1/user/whoami", AuthRequired(), whoami)
	router.GET("/api/v1/memory/:id", AuthRequired(), AuthIsOwnerOfMemoryByIdOrIsPublic(), QueryMemoryById)
	router.GET("/api/v1/favorite-location", AuthRequired(), GetMyFavLoc)
	router.GET("/api/v1/memory-my", AuthRequired(), GetPersonMemory)
	router.GET("/api/v1/user/user-info", AuthRequired(), GetUserInfo)
	router.GET("/api/v1/base-service/url", AuthRequired(), GetBaseUrl)
	router.GET("/api/v1/first-day-user", AuthRequired(), IsFirstDayRegUser)

	router.DELETE("/api/v1/memory/:id", AuthRequired(), AuthIsOwnerOfMemoryById(), DeleteMemoryById)
	router.DELETE("/api/v1/favorite-location/:id", AuthRequired(), DeleteFavoriteLocationById)

	router.PUT("/api/v1/memory/:id", AuthRequired(), AuthIsOwnerOfMemoryById(), UpdateMemoryById)
	router.PUT("/api/v1/memory-lock/:id", AuthRequired(), AuthIsOwnerOfMemoryById(), UpdateMemoryLockById)
	router.PUT("/api/v1/memory-pos/:id", AuthRequired(), UpdateMemoryPosById)
	router.PUT("/api/v1/favorite-location/used/:id", AuthRequired(), UpdateFavilocUsedTime)
	router.PUT("/api/v1/memory-content", AuthRequired(), IownMemoryFromPost, UpdateMemoryContent)

	router.POST("/login", AuthNotRequired(), login)
	router.POST("/logout", AuthRequired(), logout)
	router.POST("/api/v1/memorypoint", AuthRequired(), CreateMemoryPoint)
	router.POST("/api/v1/favorite-location", AuthRequired(), Add2FavoriteLocation)
	router.POST("/api/v1/person-memory-inbound", AuthRequired(), GetPersonMemoryInBound)
	router.POST("/api/v1/user/password", AuthRequired(), UpdatePassword)
	router.POST("/api/v1/user/readcode", AuthRequired(), UpdateReadCode)
	router.POST("/api/v1/user/user-info", AuthRequired(), UpdateUserInfo)
	router.POST("/register/email", AuthNotRequired(), RegisterByEmail)
	router.POST("/register/confirm", AuthNotRequired(), ConfirmEmail)

}
