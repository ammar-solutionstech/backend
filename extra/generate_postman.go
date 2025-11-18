package main

import (
	"encoding/json"
	"os"
)

type PostmanCollection struct {
	Info struct {
		PostmanID   string `json:"_postman_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Schema      string `json:"schema"`
	} `json:"info"`
	Item     []PostmanItem     `json:"item"`
	Variable []PostmanVariable `json:"variable"`
}

type PostmanVariable struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

type PostmanItem struct {
	Name    string          `json:"name"`
	Item    []PostmanItem   `json:"item,omitempty"`
	Request *PostmanRequest `json:"request,omitempty"`
	Event   []PostmanEvent  `json:"event,omitempty"`
}

type PostmanRequest struct {
	Auth   *PostmanAuth    `json:"auth,omitempty"`
	Method string          `json:"method"`
	Header []PostmanHeader `json:"header,omitempty"`
	Body   *PostmanBody    `json:"body,omitempty"`
	URL    PostmanURL      `json:"url"`
}

type PostmanAuth struct {
	Type   string          `json:"type"`
	Bearer []PostmanBearer `json:"bearer,omitempty"`
}

type PostmanBearer struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

type PostmanHeader struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type PostmanBody struct {
	Mode string `json:"mode"`
	Raw  string `json:"raw,omitempty"`
}

type PostmanURL struct {
	Raw      string            `json:"raw"`
	Host     []string          `json:"host"`
	Path     []string          `json:"path"`
	Variable []PostmanVariable `json:"variable,omitempty"`
}

type PostmanEvent struct {
	Listen string        `json:"listen"`
	Script PostmanScript `json:"script"`
}

type PostmanScript struct {
	Exec []string `json:"exec"`
}

func main() {
	collection := PostmanCollection{}
	collection.Info.PostmanID = "itaas-api-collection-v2"
	collection.Info.Name = "ITaaS API - Complete Collection"
	collection.Info.Description = "Complete API collection for ITaaS (IT as a Service) with all 147 endpoints"
	collection.Info.Schema = "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"

	collection.Variable = []PostmanVariable{
		{Key: "base_url", Value: "http://localhost:8080", Type: "string"},
		{Key: "token", Value: "", Type: "string"},
	}

	// Public endpoints
	publicItem := PostmanItem{
		Name: "Public",
		Item: []PostmanItem{
			createRequest("Login", "POST", "/login", false, `{"email":"user@example.com","password":"password123"}`, nil),
			createRequest("Health Check", "GET", "/health", false, "", nil),
		},
	}

	// Admin
	adminItem := PostmanItem{
		Name: "Admin",
		Item: []PostmanItem{
			createRequest("Admin Only", "GET", "/admin", true, "", nil),
		},
	}

	// Permissions (5 endpoints)
	permissionsItem := PostmanItem{
		Name: "Permissions",
		Item: []PostmanItem{
			createRequest("List All Permissions", "GET", "/api/permissions", true, "", nil),
			createRequest("Get Permission by ID", "GET", "/api/permissions/:id", true, "", []PostmanVariable{{Key: "id", Value: "1"}}),
			createRequest("Create Permission", "POST", "/api/permissions", true, `{"name":"New Permission","module":"Basic"}`, nil),
			createRequest("Update Permission", "PUT", "/api/permissions/:id", true, `{"name":"Updated Permission","module":"Basic"}`, []PostmanVariable{{Key: "id", Value: "1"}}),
			createRequest("Delete Permission", "DELETE", "/api/permissions/:id", true, "", []PostmanVariable{{Key: "id", Value: "1"}}),
		},
	}

	// Roles (8 endpoints)
	rolesItem := PostmanItem{
		Name: "Roles",
		Item: []PostmanItem{
			createRequest("List All Roles", "GET", "/api/roles", true, "", nil),
			createRequest("Get Role by ID", "GET", "/api/roles/:id", true, "", []PostmanVariable{{Key: "id", Value: "1"}}),
			createRequest("Create Role", "POST", "/api/roles", true, `{"name":"manager","description":"Manager role"}`, nil),
			createRequest("Update Role", "PUT", "/api/roles/:id", true, `{"name":"senior_manager","description":"Senior manager role"}`, []PostmanVariable{{Key: "id", Value: "1"}}),
			createRequest("Delete Role", "DELETE", "/api/roles/:id", true, "", []PostmanVariable{{Key: "id", Value: "1"}}),
			createRequest("Update Role Permissions", "PUT", "/api/roles/:id/permissions", true, `[1,2,3,4,5]`, []PostmanVariable{{Key: "id", Value: "1"}}),
			createRequest("Add Permission to Role", "POST", "/api/roles/:id/permission", true, `{"id":6}`, []PostmanVariable{{Key: "id", Value: "1"}}),
			createRequest("Delete Permission from Role", "DELETE", "/api/roles/:id/permission", true, `{"id":6}`, []PostmanVariable{{Key: "id", Value: "1"}}),
		},
	}

	// Users (12 endpoints)
	usersItem := PostmanItem{
		Name: "Users",
		Item: []PostmanItem{
			createRequest("List All Users", "GET", "/api/users", true, "", nil),
			createRequest("Get User by ID", "GET", "/api/users/:id", true, "", []PostmanVariable{{Key: "id", Value: "1"}}),
			createRequest("Create User", "POST", "/api/users", true, `{"first_name":"Jane","latest_name":"Smith","father_name":"Robert","work_email":"jane.smith@company.com","password":"password123"}`, nil),
			createRequest("Update User", "PUT", "/api/users/:id", true, `{"work_email":"new.email@company.com"}`, []PostmanVariable{{Key: "id", Value: "1"}}),
			createRequest("Delete User", "DELETE", "/api/users/:id", true, "", []PostmanVariable{{Key: "id", Value: "1"}}),
			createRequest("Suspend User", "PUT", "/api/users/:id/suspend", true, `{"active":"yes"}`, []PostmanVariable{{Key: "id", Value: "1"}}),
			createRequest("Update User Permissions", "PUT", "/api/users/:id/permissions", true, `[1,2,3,5]`, []PostmanVariable{{Key: "id", Value: "1"}}),
			createRequest("Add Permission to User", "POST", "/api/users/:id/permissions", true, `{"id":4}`, []PostmanVariable{{Key: "id", Value: "1"}}),
			createRequest("Delete Permission from User", "DELETE", "/api/users/:id/permissions", true, `{"id":4}`, []PostmanVariable{{Key: "id", Value: "1"}}),
			createRequest("Update User Roles", "PUT", "/api/users/:id/roles", true, `[1,2]`, []PostmanVariable{{Key: "id", Value: "1"}}),
			createRequest("Add Role to User", "POST", "/api/users/:id/roles", true, `{"id":3}`, []PostmanVariable{{Key: "id", Value: "1"}}),
			createRequest("Delete Role from User", "DELETE", "/api/users/:id/roles", true, `{"id":3}`, []PostmanVariable{{Key: "id", Value: "1"}}),
		},
	}

	// Help Desk Types (5 endpoints)
	helpDeskTypesItem := PostmanItem{
		Name: "Help Desk Types",
		Item: []PostmanItem{
			createRequest("List All Help Desk Types", "GET", "/api/help-desk-types", true, "", nil),
			createRequest("Get Help Desk Type by ID", "GET", "/api/help-desk-types/:id", true, "", []PostmanVariable{{Key: "id", Value: "1"}}),
			createRequest("Create Help Desk Type", "POST", "/api/help-desk-types", true, `{"name":"Software Issue","description":"Software-related problems","team_id":1,"is_active":true}`, nil),
			createRequest("Update Help Desk Type", "PUT", "/api/help-desk-types/:id", true, `{"name":"Updated Type","is_active":false}`, []PostmanVariable{{Key: "id", Value: "1"}}),
			createRequest("Delete Help Desk Type", "DELETE", "/api/help-desk-types/:id", true, "", []PostmanVariable{{Key: "id", Value: "1"}}),
		},
	}

	// Inventory - Brands (5 endpoints)
	brandsItem := PostmanItem{
		Name: "Brands",
		Item: createCRUD("Brand", "/api/brands", `{"name":"Lenovo"}`, nil),
	}

	// Inventory - Models (5 endpoints)
	modelsItem := PostmanItem{
		Name: "Models",
		Item: createCRUD("Model", "/api/models", `{"name":"ThinkPad X1 Carbon","brand_id":3}`, nil),
	}

	// Inventory - Equipment Types (5 endpoints)
	equipmentTypesItem := PostmanItem{
		Name: "Equipment Types",
		Item: createCRUD("Equipment Type", "/api/equipment-types", `{"name":"Server","description":"Server hardware"}`, nil),
	}

	// Inventory - Operating Systems (5 endpoints)
	operatingSystemsItem := PostmanItem{
		Name: "Operating Systems",
		Item: createCRUD("Operating System", "/api/operating-systems", `{"name":"Windows 11","version":"11.0"}`, nil),
	}

	// Inventory - Software Categories (5 endpoints)
	softwareCategoriesItem := PostmanItem{
		Name: "Software Categories",
		Item: createCRUD("Software Category", "/api/software-categories", `{"name":"Productivity","description":"Productivity software"}`, nil),
	}

	// Inventory - Software (5 endpoints)
	softwareItem := PostmanItem{
		Name: "Software",
		Item: createCRUD("Software", "/api/software", `{"name":"Microsoft Office","version":"2021","software_category_id":1}`, nil),
	}

	// Inventory - Equipment (5 + 11 relationship endpoints)
	equipmentSoftwareItem := PostmanItem{
		Name: "Equipment Software",
		Item: []PostmanItem{
			createRequest("List Equipment Software", "GET", "/api/equipment/:equipmentId/software", true, "", []PostmanVariable{{Key: "equipmentId", Value: "1"}}),
			createRequest("Link Software to Equipment", "POST", "/api/equipment/:equipmentId/software", true, `{"software_id":1,"install_date":"2024-01-01"}`, []PostmanVariable{{Key: "equipmentId", Value: "1"}}),
			createRequest("Update Equipment Software", "PUT", "/api/equipment/:equipmentId/software/:softwareId", true, `{"install_date":"2024-01-15"}`, []PostmanVariable{{Key: "equipmentId", Value: "1"}, {Key: "softwareId", Value: "1"}}),
			createRequest("Unlink Software from Equipment", "DELETE", "/api/equipment/:equipmentId/software/:softwareId", true, "", []PostmanVariable{{Key: "equipmentId", Value: "1"}, {Key: "softwareId", Value: "1"}}),
		},
	}

	equipmentHelpDeskItem := PostmanItem{
		Name: "Equipment Help Desk",
		Item: []PostmanItem{
			createRequest("List Equipment Help Desk Links", "GET", "/api/equipment/:equipmentId/help-desk", true, "", []PostmanVariable{{Key: "equipmentId", Value: "1"}}),
			createRequest("Link Equipment to Help Desk", "POST", "/api/equipment/:equipmentId/help-desk", true, `{"help_desk_id":1}`, []PostmanVariable{{Key: "equipmentId", Value: "1"}}),
			createRequest("Unlink Equipment from Help Desk", "DELETE", "/api/equipment/:equipmentId/help-desk/:helpDeskId", true, "", []PostmanVariable{{Key: "equipmentId", Value: "1"}, {Key: "helpDeskId", Value: "1"}}),
		},
	}

	equipmentUserHistoryItem := PostmanItem{
		Name: "Equipment User History",
		Item: []PostmanItem{
			createRequest("List Equipment User History", "GET", "/api/equipment/:equipmentId/user-history", true, "", []PostmanVariable{{Key: "equipmentId", Value: "1"}}),
			createRequest("Create Equipment User History", "POST", "/api/equipment/:equipmentId/user-history", true, `{"user_id":1,"start_date":"2024-01-01","end_date":"2024-12-31"}`, []PostmanVariable{{Key: "equipmentId", Value: "1"}}),
			createRequest("Update Equipment User History", "PUT", "/api/equipment/:equipmentId/user-history/:userId/:startDate", true, `{"end_date":"2024-06-30"}`, []PostmanVariable{{Key: "equipmentId", Value: "1"}, {Key: "userId", Value: "1"}, {Key: "startDate", Value: "2024-01-01"}}),
			createRequest("Delete Equipment User History", "DELETE", "/api/equipment/:equipmentId/user-history/:userId/:startDate", true, "", []PostmanVariable{{Key: "equipmentId", Value: "1"}, {Key: "userId", Value: "1"}, {Key: "startDate", Value: "2024-01-01"}}),
		},
	}

	equipmentCRUD := createCRUD("Equipment", "/api/equipment", `{"name":"Laptop-001","description":"Dell Latitude","serial_number":"SN123","equipment_type_id":1,"model_id":1,"production_date":"2024-01-01","ip_v4":"192.168.1.1","ip_v6":"::1","mac_address":"00:11:22:33:44:55","operating_system_id":1,"location_id":1,"proccessors":"Intel i7","country_of_region":1,"user_id":1,"warrantly_start_date":"2024-01-01","warrantly_end_date":"2027-01-01","supplier_id":1}`, nil)
	equipmentItem := PostmanItem{
		Name: "Equipment",
		Item: append(append(append(equipmentCRUD, equipmentSoftwareItem), equipmentHelpDeskItem), equipmentUserHistoryItem),
	}

	// Inventory - Documents (5 endpoints)
	documentsItem := PostmanItem{
		Name: "Documents",
		Item: createCRUD("Document", "/api/documents", `{"name":"Warranty.pdf","document_size":1024,"document_type":"pdf","equipment_id":1}`, nil),
	}

	// Inventory - Maintenance (5 endpoints)
	maintenanceItem := PostmanItem{
		Name: "Maintenance",
		Item: createCRUD("Maintenance Record", "/api/maintenance", `{"equipment_id":1,"maintenance_date":"2024-01-15","description":"Regular maintenance","cost":100.50}`, nil),
	}

	// Geography - Countries (5 endpoints)
	countriesItem := PostmanItem{
		Name: "Countries",
		Item: createCRUD("Country", "/api/countries", `{"name":"United States","code":"US"}`, nil),
	}

	// Geography - Cities (5 endpoints)
	citiesItem := PostmanItem{
		Name: "Cities",
		Item: createCRUD("City", "/api/cities", `{"name":"New York","country_id":1}`, nil),
	}

	// Geography - Locations (5 endpoints)
	locationsItem := PostmanItem{
		Name: "Locations",
		Item: createCRUD("Location", "/api/locations", `{"name":"Head Office","address":"123 Main St","city_id":1}`, nil),
	}

	// Geography - Contacts (5 endpoints)
	contactsItem := PostmanItem{
		Name: "Contacts",
		Item: createCRUD("Contact", "/api/contacts", `{"name":"ABC Corp","email":"contact@abc.com","mobile_number":"+1234567890","country_id":1}`, nil),
	}

	// Help Desk - Tickets (5 + 3 relationship endpoints)
	helpDeskParticipantsItem := PostmanItem{
		Name: "Participants",
		Item: []PostmanItem{
			createRequest("List Help Desk Participants", "GET", "/api/help-desk/:helpDeskId/participants", true, "", []PostmanVariable{{Key: "helpDeskId", Value: "1"}}),
			createRequest("Add Help Desk Participant", "POST", "/api/help-desk/:helpDeskId/participants", true, `{"user_id":1}`, []PostmanVariable{{Key: "helpDeskId", Value: "1"}}),
			createRequest("Remove Help Desk Participant", "DELETE", "/api/help-desk/:helpDeskId/participants/:userId", true, "", []PostmanVariable{{Key: "helpDeskId", Value: "1"}, {Key: "userId", Value: "1"}}),
		},
	}

	helpDeskTicketsCRUD := createCRUD("Help Desk Ticket", "/api/help-desk", `{"name":"Laptop won't start","create_date":"2024-01-15","help_desk_type_id":1,"portal_user_id":1,"description":"Black screen on boot","state":"open"}`, nil)
	helpDeskTicketsItem := PostmanItem{
		Name: "Tickets",
		Item: append(helpDeskTicketsCRUD, helpDeskParticipantsItem),
	}

	// Help Desk - Ratings (5 endpoints)
	ratingsItem := PostmanItem{
		Name: "Ratings",
		Item: createCRUD("Help Desk Rating", "/api/help-desk/ratings", `{"help_desk_id":1,"rating":5,"comment":"Excellent service"}`, nil),
	}

	// Help Desk - Transactions (5 + 3 relationship endpoints)
	transactionUsersItem := PostmanItem{
		Name: "Transaction Users",
		Item: []PostmanItem{
			createRequest("List Transaction Users", "GET", "/api/help-desk/transactions/:transactionId/users", true, "", []PostmanVariable{{Key: "transactionId", Value: "1"}}),
			createRequest("Add Transaction User", "POST", "/api/help-desk/transactions/:transactionId/users", true, `{"user_id":1}`, []PostmanVariable{{Key: "transactionId", Value: "1"}}),
			createRequest("Remove Transaction User", "DELETE", "/api/help-desk/transactions/:transactionId/users/:userId", true, "", []PostmanVariable{{Key: "transactionId", Value: "1"}, {Key: "userId", Value: "1"}}),
		},
	}

	transactionsCRUD := createCRUD("Help Desk Transaction", "/api/help-desk/transactions", `{"help_desk_id":1,"transaction_type_id":1,"description":"Initial response","date":"2024-01-15"}`, nil)
	transactionsItem := PostmanItem{
		Name: "Transactions",
		Item: append(transactionsCRUD, transactionUsersItem),
	}

	// Help Desk - Teams (5 + 3 relationship endpoints)
	teamMembersItem := PostmanItem{
		Name: "Team Members",
		Item: []PostmanItem{
			createRequest("List Team Members", "GET", "/api/help-desk/teams/:teamId/members", true, "", []PostmanVariable{{Key: "teamId", Value: "1"}}),
			createRequest("Add Team Member", "POST", "/api/help-desk/teams/:teamId/members", true, `{"user_id":1}`, []PostmanVariable{{Key: "teamId", Value: "1"}}),
			createRequest("Remove Team Member", "DELETE", "/api/help-desk/teams/:teamId/members/:userId", true, "", []PostmanVariable{{Key: "teamId", Value: "1"}, {Key: "userId", Value: "1"}}),
		},
	}

	teamsCRUD := createCRUD("Help Desk Team", "/api/help-desk/teams", `{"name":"IT Support","description":"Primary IT support team","manager_id":1,"is_active":true}`, nil)
	teamsItem := PostmanItem{
		Name: "Teams",
		Item: append(teamsCRUD, teamMembersItem),
	}

	// Help Desk - Transaction Types (5 endpoints)
	transactionTypesItem := PostmanItem{
		Name: "Transaction Types",
		Item: createCRUD("Help Desk Transaction Type", "/api/help-desk/transaction-types", `{"name":"Response","description":"Initial response to ticket"}`, nil),
	}

	// Navigation - Menus (5 + 3 relationship endpoints)
	menuRolesItem := PostmanItem{
		Name: "Menu Roles",
		Item: []PostmanItem{
			createRequest("List Menu Roles", "GET", "/api/menus/:menuId/roles", true, "", []PostmanVariable{{Key: "menuId", Value: "1"}}),
			createRequest("Add Menu Role", "POST", "/api/menus/:menuId/roles", true, `{"role_id":1}`, []PostmanVariable{{Key: "menuId", Value: "1"}}),
			createRequest("Remove Menu Role", "DELETE", "/api/menus/:menuId/roles/:roleId", true, "", []PostmanVariable{{Key: "menuId", Value: "1"}, {Key: "roleId", Value: "1"}}),
		},
	}

	menusCRUD := createCRUD("Menu", "/api/menus", `{"name":"Dashboard","path":"/dashboard","icon":"home","parent_id":null,"order":1}`, nil)
	menusItem := PostmanItem{
		Name: "Menus",
		Item: append(menusCRUD, menuRolesItem),
	}

	// Assemble collection
	collection.Item = []PostmanItem{
		publicItem,
		adminItem,
		permissionsItem,
		rolesItem,
		usersItem,
		helpDeskTypesItem,
		{
			Name: "Inventory",
			Item: []PostmanItem{
				brandsItem, modelsItem, equipmentTypesItem, operatingSystemsItem,
				softwareCategoriesItem, softwareItem, equipmentItem, documentsItem, maintenanceItem,
			},
		},
		{
			Name: "Geography",
			Item: []PostmanItem{countriesItem, citiesItem, locationsItem, contactsItem},
		},
		{
			Name: "Help Desk",
			Item: []PostmanItem{helpDeskTicketsItem, ratingsItem, transactionsItem, teamsItem, transactionTypesItem},
		},
		{
			Name: "Navigation",
			Item: []PostmanItem{menusItem},
		},
	}

	// Add login event to save token
	loginReq := collection.Item[0].Item[0].Request
	if loginReq != nil {
		collection.Item[0].Item[0].Event = []PostmanEvent{
			{
				Listen: "test",
				Script: PostmanScript{
					Exec: []string{
						"if (pm.response.code === 200) {",
						"    var jsonData = pm.response.json();",
						"    pm.environment.set(\"token\", jsonData.token);",
						"}",
					},
				},
			},
		}
	}

	// Write JSON
	jsonData, err := json.MarshalIndent(collection, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("ITaaS_API.postman_collection.json", jsonData, 0644)
	if err != nil {
		panic(err)
	}
}

func createRequest(name, method, path string, requiresAuth bool, body string, vars []PostmanVariable) PostmanItem {
	item := PostmanItem{Name: name}

	parts := []string{}
	for _, p := range splitPath(path) {
		if p != "" {
			parts = append(parts, p)
		}
	}

	url := PostmanURL{
		Raw:  "{{base_url}}" + path,
		Host: []string{"{{base_url}}"},
		Path: parts,
	}
	if vars != nil {
		url.Variable = vars
	}

	req := PostmanRequest{
		Method: method,
		URL:    url,
	}

	if requiresAuth {
		req.Auth = &PostmanAuth{
			Type: "bearer",
			Bearer: []PostmanBearer{
				{Key: "token", Value: "{{token}}", Type: "string"},
			},
		}
	}

	if body != "" {
		req.Header = []PostmanHeader{
			{Key: "Content-Type", Value: "application/json"},
		}
		req.Body = &PostmanBody{
			Mode: "raw",
			Raw:  body,
		}
	}

	item.Request = &req
	return item
}

func createCRUD(resourceName, basePath string, createBody string, vars []PostmanVariable) []PostmanItem {
	return []PostmanItem{
		createRequest("List "+resourceName+"s", "GET", basePath, true, "", nil),
		createRequest("Get "+resourceName, "GET", basePath+"/:id", true, "", append(vars, PostmanVariable{Key: "id", Value: "1"})),
		createRequest("Create "+resourceName, "POST", basePath, true, createBody, vars),
		createRequest("Update "+resourceName, "PUT", basePath+"/:id", true, createBody, append(vars, PostmanVariable{Key: "id", Value: "1"})),
		createRequest("Delete "+resourceName, "DELETE", basePath+"/:id", true, "", append(vars, PostmanVariable{Key: "id", Value: "1"})),
	}
}

func splitPath(path string) []string {
	parts := []string{}
	current := ""
	for _, char := range path {
		if char == '/' {
			if current != "" {
				parts = append(parts, current)
				current = ""
			}
			parts = append(parts, "")
		} else {
			current += string(char)
		}
	}
	if current != "" {
		parts = append(parts, current)
	}
	return parts
}
