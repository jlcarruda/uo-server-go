package UoServer

import uuid "github.com/nu7hatch/gouuid"

func ClientFilter(clients []*Client, f func(client *Client) bool) []*Client {
	filtered := make([]*Client, 0)

	for _, c := range clients {
		if f(c) {
			filtered = append(filtered, c)
		}
	}

	return filtered
}

func RemoveClientByID(id *uuid.UUID) {
	filtered := ClientFilter(CLIENTS, func (client *Client) bool {
		return client.id != id
	})

	CLIENTS = filtered
}

func GetClientByID(id *uuid.UUID) *Client {
	filtered := ClientFilter(CLIENTS, func (client *Client) bool {
		return client.id == id
	})

	return filtered[0]
}

