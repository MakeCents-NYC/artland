package mastercard

type AppConfig struct {
	Definition  AppInfoDefinition
	Description *string // Description of your blockchain application
	Name        string  // Your application identifier. MUST be the protobuf package name
	Version     *string // Version of your blockchain application
}

type AppInfo struct {
	Definition *AppInfoDefinition
	Id         *string // Your application identifier. MUST be the protobuf package name
}

type AppInfoDefinition struct {
	Encoding string // Encoding scheme used to encode your protocol. base58, base64, hex
	Format   string // must be "proto3"
	Message  string // Your protocol buffer message definition
}

type AuthorizationRequest struct {
	Amount_min_or_units int    // The amount to authorize, in minor units of the specified currency
	Currency            string // ISO 4217 Currency Code specifying the currency
	Description         string // A description of the event for which the authorization message is to be generated
	Merchant_id         string // Unique identifier of the merchant
	Pan                 string // The Personal Account Number of the card for which to generate an authorization message
	Pos_id              string // An identifier for the point-of-sale terminal for which the authorization message is to be generated
}

type AuthorizationResponse struct {
	Encoding   string // Encoding scheme used to encode your protocol. base58, base64, hex
	Public_key string // Public key in which the signature may be verified
	Reference  string // A unique reference to an authorization event
}

type Block struct {
	Authority      *string // The signing identity from the audit nodes, either as a blockchain address, or as the raw public key
	Hash           *string // The hash of this block (hex)
	Nonce          *string // Random number used in the generation of this block
	Partitions     []DataPartition
	Previous_Block *string // hash of the previous block (hex)
	Signature      *string // The signature confirmation from the audit nodes
	Slot           *int32  // The slot number of this block
	Version        *string // The version of the block format, 1.0.0.0
}

type Connection struct {
	Address *string // The address of the connected node (base58)
}

type CreatedNode struct {
	Address   *string      // The address of your newly created blockchain node
	Authority *string      // The signing identity from the audit nodes, either as a blockchain address, or as the raw public key
	Peers     []Connection // Connected nodes
	Type      *string      // customer
}

type CurrentBlock struct {
	Ref  *string // Hash of the last confirmed block (hex)
	Slot *int32  // Slot number of the last confirmed block
}

type DataPartition struct {
	Application *string  // Application ID
	Entries     []string // List of hashes for all entries in this partition
	Entry_Count *int     // Number of entries referred to by the merkle root hash
	Merkle_Root *string  // Its the root hash of all entries for this application in this block (hex)
}

type EntryInfo struct {
	Hash   *string // (hex)
	Slot   *int    // Slot this entry is written to
	Status *string // Enum pending, confirmed, rejected
	Value  *string // Hex representation of the value written onto the chain
}
type EntryRequest struct {
	App      string // The application ID
	Encoding string // Enum base58, base64, hex
	Value    string // Value to be written onto the chain. Should be encoded according to above
}

type EntryResponse struct {
	Hash   *string // Hash of the value written onto the chain. Can be used to lookup the entry. (hex)
	Slot   *int    // Slot the entry is written to. Can be used to query a block.
	Status *string // Enum pending, confirmed, rejected
}

type GenesisBlock struct {
	Ref  *string // Hash of the first block on this network
	Slot *int32  // The slot number of the first block chain on this network
}

type InviteRequest struct {
	Invitations []string // List of email addresses
}

type InviteResponse struct {
	Tokens string // An array of invitation/activation codes (hex)
}

type JoinRequest struct {
	Activation_key string // Invitation code you received in your MasterCard blockchain invitation email
}

type NodeDefinition struct {
	Application *AppConfig
	Invitations []string // List of email addresses of others you would like to invite to your BlockChain application
	Network     *string  // Use zone. Defaults to "ZONE"
}

type NodeInfo struct {
	Address      *string      // The address of this node (base58)
	Authority    *string      // The signing identity from the audit nodes, either as a blockchain address, or as the raw public key (base58)
	Chain_Height *int32       // Last verified slot on the chain as seen by this node
	Delay        *int         // Consensus delay - expected delay from when you send a transaction entry and when it is expected to be written into a block on the chain
	Drift        *int         // Time drift on this node
	Peers        []Connection // Connected nodes
	Public_Key   *string      // Public key of node. Can be used to validate signatures from this node
	Type         *string      // Enum audit, consensus, application, customer
	Unconfirmed  *int32       // The number of blockchain entries awaiting confirmation
}

type SettlementRequest struct {
	Encoding string // Encoding scheme used to encode your protocol. base58, base64, hex
	Hash     string // A hash reference to the blockchain entry in which the settlement request occurs on the chain
}

type SettlementResponse struct {
	Encoding   string // Encoding scheme used to encode your protocol. base58, base64, hex
	Public_Key string // secp2561 public key with which the signature in a settlement notification on the blockchain may be verified
	Signature  string // The signed hash of the request using the private key corresponding to the public key above
}

type SignRequest struct {
	Encoding string   // Encoding scheme used to encode your protocol. base58, base64, hex
	Values   []string // An array of values
}

type SignResponse struct {
	Encoding   string   // Encoding scheme used to encode your protocol. base58, base64, hex
	Signatures []string // An array of ECDSA signatures
}

type StatusResponse struct {
	Alerts       []string // Informational alerts such as upcoming version or config changes
	Applications []string // The application identifiers that are on this node
	Current      *CurrentBlock
	Genesis      *GenesisBlock
	Network      int //Should always be 1513115205
}

type ValidationError struct {
	Code    *int32
	Fields  []string
	Message *string
}

type SimpleError struct {
	Code    *int32
	Message *string
}
