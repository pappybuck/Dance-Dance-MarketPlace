-- CreateTable
CREATE TABLE "Product" (
    "id" UUID NOT NULL DEFAULT gen_random_uuid(),
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "price"  numeric(65, 30) NOT NULL,
    "quantity" INTEGER NOT NULL,

    CONSTRAINT "Product_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Review" (
    "id" UUID NOT NULL DEFAULT gen_random_uuid(),
    "productId" UUID NOT NULL,
    "profileId" UUID NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "rating" INTEGER NOT NULL,

    CONSTRAINT "Review_pkey" PRIMARY KEY ("id")
);

CREATE TABLE "Profile" (
    "id" UUID NOT NULL DEFAULT gen_random_uuid(),
    "userId" UUID NOT NULL,
    "name" TEXT NOT NULL,
    "phone" TEXT NOT NULL DEFAULT 'N/A',
    CONSTRAINT "Profile_pkey" PRIMARY KEY ("id")
);

CREATE TABLE "Role" (
    "id" SERIAL NOT NULL,
    "name" TEXT NOT NULL,

    CONSTRAINT "Role_pkey" PRIMARY KEY ("id")
);


-- CreateTable
CREATE TABLE "User" (
    "id" UUID NOT NULL DEFAULT gen_random_uuid(),
    "roleId" INTEGER NOT NULL DEFAULT 2,
    "email" TEXT NOT NULL,
    "password" TEXT NOT NULL,
    CONSTRAINT "User_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "User_email_key" ON "User"("email");

-- AddForeignKey
ALTER TABLE "Review" ADD CONSTRAINT "Review_productId_fkey" FOREIGN KEY ("productId") REFERENCES "Product"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Review" ADD CONSTRAINT "Review_profileId_fkey" FOREIGN KEY ("profileId") REFERENCES "Profile"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Profile" ADD CONSTRAINT "Profile_userId_fkey" FOREIGN KEY ("userId") REFERENCES "User"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "User" ADD CONSTRAINT "Profile_roleId_fkey" FOREIGN KEY ("roleId") REFERENCES "Role"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- Seed Data
INSERT INTO "Products" ("id", "name", "description", "price", "quantity") VALUES ('84726401-aa95-4d1e-a30d-cd7642528d6f', 'Product 1', 'Description 1', 1.00, 1);
INSERT INTO "Products" ("id", "name", "description", "price", "quantity") VALUES ('8a6425b9-4127-4856-9aa7-0b001c6cc773', 'Product 2', 'Description 2', 2.00, 2);
INSERT INTO "Products" ("id", "name", "description", "price", "quantity") VALUES ('a923516e-3363-4940-8d54-d426a9655447', 'Product 3', 'Description 3', 3.00, 3);
INSERT INTO "Products" ("id", "name", "description", "price", "quantity") VALUES ('f31e656f-5b23-4241-b7ea-2345ffe1041e', 'Product 4', 'Description 4', 4.00, 4);
INSERT INTO "Products" ("id", "name", "description", "price", "quantity") VALUES ('8b572ea7-1bd2-4a87-9922-8a161c3238b8', 'Product 5', 'Description 5', 5.00, 5);
INSERT INTO "Products" ("id", "name", "description", "price", "quantity") VALUES ('af7ffbd9-b761-492e-8bcf-fced78ab593a', 'Product 6', 'Description 6', 6.00, 6);
INSERT INTO "Products" ("id", "name", "description", "price", "quantity") VALUES ('c1526900-efe1-4830-9d18-63772453b3ad', 'Product 7', 'Description 7', 7.00, 7);
INSERT INTO "Products" ("id", "name", "description", "price", "quantity") VALUES ('7e3f1689-8af1-40ff-98bc-b9e03ca597ea', 'Product 8', 'Description 8', 8.00, 8);
INSERT INTO "Products" ("id", "name", "description", "price", "quantity") VALUES ('76f232c9-06a3-4788-8cb9-ea4876db1edb', 'Product 9', 'Description 9', 9.00, 9);
INSERT INTO "Products" ("id", "name", "description", "price", "quantity") VALUES ('89aad33c-c1e4-4605-aa2b-0b2056c138f0', 'Product 10', 'Description 10', 10.00, 10);
INSERT INTO "Role" ("id", "name") VALUES (1, 'Admin');
INSERT INTO "Role" ("id", "name") VALUES (2, 'User');
INSERT INTO "User" ("id", "email", "password") VALUES ('3788c60a-b9fe-4bc4-b7a7-5d0df7bdf5c5', 'email 1', 'password 1');
INSERT INTO "Profile" ("id", "userId", "name") VALUES ('39744f89-6e5f-4fbd-9377-3ca3f9c6a729', '3788c60a-b9fe-4bc4-b7a7-5d0df7bdf5c5', 'Profile 1');
INSERT INTO "User" ("id", "email", "password") VALUES ('73189c14-0e50-4525-9c5d-31cb1aa3057b', 'email 2', 'password 2');
INSERT INTO "Profile" ("id", "userId", "name") VALUES ('d9319d51-613b-452d-a3c2-be2e8c4e8df2', '73189c14-0e50-4525-9c5d-31cb1aa3057b', 'Profile 2');
INSERT INTO "User" ("id", "email", "password") VALUES ('3c14e39b-5acd-41eb-9a8a-415f4b64106e', 'email 3', 'password 3');
INSERT INTO "Profile" ("id", "userId", "name") VALUES ('829ae5be-0558-49a2-a0aa-3fe925e9faba', '3c14e39b-5acd-41eb-9a8a-415f4b64106e', 'Profile 3');
INSERT INTO "User" ("id", "email", "password") VALUES ('0bbd55f6-95a6-4197-9a6c-34e2bd7b7c77', 'email 4', 'password 4');
INSERT INTO "Profile" ("id", "userId", "name") VALUES ('562a9c34-594f-48fb-a4ff-9540aef4abd0', '0bbd55f6-95a6-4197-9a6c-34e2bd7b7c77', 'Profile 4');
INSERT INTO "User" ("id", "email", "password") VALUES ('412b3a7c-4c6d-42a8-8010-de966c37cdad', 'email 5', 'password 5');
INSERT INTO "Profile" ("id", "userId", "name") VALUES ('6f8e665a-a9cf-43e8-92e0-6a2c70feea77', '412b3a7c-4c6d-42a8-8010-de966c37cdad', 'Profile 5');
INSERT INTO "Review" ("productId", "profileId", "name", "description", "rating") VALUES ('84726401-aa95-4d1e-a30d-cd7642528d6f', '39744f89-6e5f-4fbd-9377-3ca3f9c6a729', 'Review 1', 'Description 1', 1);
INSERT INTO "Review" ("productId", "profileId", "name", "description", "rating") VALUES ('84726401-aa95-4d1e-a30d-cd7642528d6f', 'd9319d51-613b-452d-a3c2-be2e8c4e8df2', 'Review 2', 'Description 2', 2);
INSERT INTO "Review" ("productId", "profileId", "name", "description", "rating") VALUES ('84726401-aa95-4d1e-a30d-cd7642528d6f', '829ae5be-0558-49a2-a0aa-3fe925e9faba', 'Review 3', 'Description 3', 3);
INSERT INTO "Review" ("productId", "profileId", "name", "description", "rating") VALUES ('84726401-aa95-4d1e-a30d-cd7642528d6f', '562a9c34-594f-48fb-a4ff-9540aef4abd0', 'Review 4', 'Description 4', 4);
INSERT INTO "Review" ("productId", "profileId", "name", "description", "rating") VALUES ('84726401-aa95-4d1e-a30d-cd7642528d6f', '6f8e665a-a9cf-43e8-92e0-6a2c70feea77', 'Review 5', 'Description 5', 5);
INSERT INTO "Review" ("productId", "profileId", "name", "description", "rating") VALUES ('8a6425b9-4127-4856-9aa7-0b001c6cc773', '39744f89-6e5f-4fbd-9377-3ca3f9c6a729', 'Review 6', 'Description 6', 6);
INSERT INTO "Review" ("productId", "profileId", "name", "description", "rating") VALUES ('8a6425b9-4127-4856-9aa7-0b001c6cc773', 'd9319d51-613b-452d-a3c2-be2e8c4e8df2', 'Review 7', 'Description 7', 7);
INSERT INTO "Review" ("productId", "profileId", "name", "description", "rating") VALUES ('8a6425b9-4127-4856-9aa7-0b001c6cc773', '829ae5be-0558-49a2-a0aa-3fe925e9faba', 'Review 8', 'Description 8', 8);
INSERT INTO "Review" ("productId", "profileId", "name", "description", "rating") VALUES ('8a6425b9-4127-4856-9aa7-0b001c6cc773', '562a9c34-594f-48fb-a4ff-9540aef4abd0', 'Review 9', 'Description 9', 9);
INSERT INTO "Review" ("productId", "profileId", "name", "description", "rating") VALUES ('8a6425b9-4127-4856-9aa7-0b001c6cc773', '6f8e665a-a9cf-43e8-92e0-6a2c70feea77', 'Review 10', 'Description 10', 10);
INSERT INTO "Review" ("productId", "profileId", "name", "description", "rating") VALUES ('a923516e-3363-4940-8d54-d426a9655447', '39744f89-6e5f-4fbd-9377-3ca3f9c6a729', 'Review 11', 'Description 11', 11);
INSERT INTO "Review" ("productId", "profileId", "name", "description", "rating") VALUES ('a923516e-3363-4940-8d54-d426a9655447', 'd9319d51-613b-452d-a3c2-be2e8c4e8df2', 'Review 12', 'Description 12', 12);
INSERT INTO "Review" ("productId", "profileId", "name", "description", "rating") VALUES ('a923516e-3363-4940-8d54-d426a9655447', '829ae5be-0558-49a2-a0aa-3fe925e9faba', 'Review 13', 'Description 13', 13);
INSERT INTO "Review" ("productId", "profileId", "name", "description", "rating") VALUES ('a923516e-3363-4940-8d54-d426a9655447', '562a9c34-594f-48fb-a4ff-9540aef4abd0', 'Review 14', 'Description 14', 14);
INSERT INTO "Review" ("productId", "profileId", "name", "description", "rating") VALUES ('a923516e-3363-4940-8d54-d426a9655447', '6f8e665a-a9cf-43e8-92e0-6a2c70feea77', 'Review 15', 'Description 15', 15);
INSERT INTO "Review" ("productId", "profileId", "name", "description", "rating") VALUES ('f31e656f-5b23-4241-b7ea-2345ffe1041e', '39744f89-6e5f-4fbd-9377-3ca3f9c6a729', 'Review 16', 'Description 16', 16);
INSERT INTO "Review" ("productId", "profileId", "name", "description", "rating") VALUES ('f31e656f-5b23-4241-b7ea-2345ffe1041e', 'd9319d51-613b-452d-a3c2-be2e8c4e8df2', 'Review 17', 'Description 17', 17);
INSERT INTO "Review" ("productId", "profileId", "name", "description", "rating") VALUES ('f31e656f-5b23-4241-b7ea-2345ffe1041e', '829ae5be-0558-49a2-a0aa-3fe925e9faba', 'Review 18', 'Description 18', 18);
INSERT INTO "Review" ("productId", "profileId", "name", "description", "rating") VALUES ('f31e656f-5b23-4241-b7ea-2345ffe1041e', '562a9c34-594f-48fb-a4ff-9540aef4abd0', 'Review 19', 'Description 19', 19);