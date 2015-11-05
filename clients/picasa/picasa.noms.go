// This file was generated by nomdl/codegen.

package main

import (
	"github.com/attic-labs/noms/chunks"
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __mainPackageInFile_picasa_CachedRef ref.Ref

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func init() {
	p := types.NewPackage([]types.TypeRef{
		types.MakeStructTypeRef("User",
			[]types.Field{
				types.Field{"Id", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"Name", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"Albums", types.MakeCompoundTypeRef(types.MapKind, types.MakePrimitiveTypeRef(types.StringKind), types.MakeTypeRef(ref.Ref{}, 1)), false},
				types.Field{"RefreshToken", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"OAuthToken", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"OAuthSecret", types.MakePrimitiveTypeRef(types.StringKind), false},
			},
			types.Choices{},
		),
		types.MakeStructTypeRef("Album",
			[]types.Field{
				types.Field{"Id", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"Title", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"NumPhotos", types.MakePrimitiveTypeRef(types.UInt32Kind), false},
				types.Field{"Photos", types.MakeCompoundTypeRef(types.RefKind, types.MakeCompoundTypeRef(types.SetKind, types.MakeCompoundTypeRef(types.RefKind, types.MakeTypeRef(ref.Parse("sha1-00419ebbb418539af67238164b20341913efeb4d"), 0)))), false},
			},
			types.Choices{},
		),
	}, []ref.Ref{
		ref.Parse("sha1-00419ebbb418539af67238164b20341913efeb4d"),
	})
	__mainPackageInFile_picasa_CachedRef = types.RegisterPackage(&p)
}

// User

type User struct {
	_Id           string
	_Name         string
	_Albums       MapOfStringToAlbum
	_RefreshToken string
	_OAuthToken   string
	_OAuthSecret  string

	ref *ref.Ref
}

func NewUser() User {
	return User{
		_Id:           "",
		_Name:         "",
		_Albums:       NewMapOfStringToAlbum(),
		_RefreshToken: "",
		_OAuthToken:   "",
		_OAuthSecret:  "",

		ref: &ref.Ref{},
	}
}

type UserDef struct {
	Id           string
	Name         string
	Albums       MapOfStringToAlbumDef
	RefreshToken string
	OAuthToken   string
	OAuthSecret  string
}

func (def UserDef) New() User {
	return User{
		_Id:           def.Id,
		_Name:         def.Name,
		_Albums:       def.Albums.New(),
		_RefreshToken: def.RefreshToken,
		_OAuthToken:   def.OAuthToken,
		_OAuthSecret:  def.OAuthSecret,
		ref:           &ref.Ref{},
	}
}

func (s User) Def() (d UserDef) {
	d.Id = s._Id
	d.Name = s._Name
	d.Albums = s._Albums.Def()
	d.RefreshToken = s._RefreshToken
	d.OAuthToken = s._OAuthToken
	d.OAuthSecret = s._OAuthSecret
	return
}

var __typeRefForUser types.TypeRef

func (m User) TypeRef() types.TypeRef {
	return __typeRefForUser
}

func init() {
	__typeRefForUser = types.MakeTypeRef(__mainPackageInFile_picasa_CachedRef, 0)
	types.RegisterStruct(__typeRefForUser, builderForUser, readerForUser)
}

func builderForUser(values []types.Value) types.Value {
	i := 0
	s := User{ref: &ref.Ref{}}
	s._Id = values[i].(types.String).String()
	i++
	s._Name = values[i].(types.String).String()
	i++
	s._Albums = values[i].(MapOfStringToAlbum)
	i++
	s._RefreshToken = values[i].(types.String).String()
	i++
	s._OAuthToken = values[i].(types.String).String()
	i++
	s._OAuthSecret = values[i].(types.String).String()
	i++
	return s
}

func readerForUser(v types.Value) []types.Value {
	values := []types.Value{}
	s := v.(User)
	values = append(values, types.NewString(s._Id))
	values = append(values, types.NewString(s._Name))
	values = append(values, s._Albums)
	values = append(values, types.NewString(s._RefreshToken))
	values = append(values, types.NewString(s._OAuthToken))
	values = append(values, types.NewString(s._OAuthSecret))
	return values
}

func (s User) Equals(other types.Value) bool {
	return other != nil && __typeRefForUser.Equals(other.TypeRef()) && s.Ref() == other.Ref()
}

func (s User) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s User) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeRefForUser.Chunks()...)
	chunks = append(chunks, s._Albums.Chunks()...)
	return
}

func (s User) Id() string {
	return s._Id
}

func (s User) SetId(val string) User {
	s._Id = val
	s.ref = &ref.Ref{}
	return s
}

func (s User) Name() string {
	return s._Name
}

func (s User) SetName(val string) User {
	s._Name = val
	s.ref = &ref.Ref{}
	return s
}

func (s User) Albums() MapOfStringToAlbum {
	return s._Albums
}

func (s User) SetAlbums(val MapOfStringToAlbum) User {
	s._Albums = val
	s.ref = &ref.Ref{}
	return s
}

func (s User) RefreshToken() string {
	return s._RefreshToken
}

func (s User) SetRefreshToken(val string) User {
	s._RefreshToken = val
	s.ref = &ref.Ref{}
	return s
}

func (s User) OAuthToken() string {
	return s._OAuthToken
}

func (s User) SetOAuthToken(val string) User {
	s._OAuthToken = val
	s.ref = &ref.Ref{}
	return s
}

func (s User) OAuthSecret() string {
	return s._OAuthSecret
}

func (s User) SetOAuthSecret(val string) User {
	s._OAuthSecret = val
	s.ref = &ref.Ref{}
	return s
}

// Album

type Album struct {
	_Id        string
	_Title     string
	_NumPhotos uint32
	_Photos    RefOfSetOfRefOfRemotePhoto

	ref *ref.Ref
}

func NewAlbum() Album {
	return Album{
		_Id:        "",
		_Title:     "",
		_NumPhotos: uint32(0),
		_Photos:    NewRefOfSetOfRefOfRemotePhoto(ref.Ref{}),

		ref: &ref.Ref{},
	}
}

type AlbumDef struct {
	Id        string
	Title     string
	NumPhotos uint32
	Photos    ref.Ref
}

func (def AlbumDef) New() Album {
	return Album{
		_Id:        def.Id,
		_Title:     def.Title,
		_NumPhotos: def.NumPhotos,
		_Photos:    NewRefOfSetOfRefOfRemotePhoto(def.Photos),
		ref:        &ref.Ref{},
	}
}

func (s Album) Def() (d AlbumDef) {
	d.Id = s._Id
	d.Title = s._Title
	d.NumPhotos = s._NumPhotos
	d.Photos = s._Photos.TargetRef()
	return
}

var __typeRefForAlbum types.TypeRef

func (m Album) TypeRef() types.TypeRef {
	return __typeRefForAlbum
}

func init() {
	__typeRefForAlbum = types.MakeTypeRef(__mainPackageInFile_picasa_CachedRef, 1)
	types.RegisterStruct(__typeRefForAlbum, builderForAlbum, readerForAlbum)
}

func builderForAlbum(values []types.Value) types.Value {
	i := 0
	s := Album{ref: &ref.Ref{}}
	s._Id = values[i].(types.String).String()
	i++
	s._Title = values[i].(types.String).String()
	i++
	s._NumPhotos = uint32(values[i].(types.UInt32))
	i++
	s._Photos = values[i].(RefOfSetOfRefOfRemotePhoto)
	i++
	return s
}

func readerForAlbum(v types.Value) []types.Value {
	values := []types.Value{}
	s := v.(Album)
	values = append(values, types.NewString(s._Id))
	values = append(values, types.NewString(s._Title))
	values = append(values, types.UInt32(s._NumPhotos))
	values = append(values, s._Photos)
	return values
}

func (s Album) Equals(other types.Value) bool {
	return other != nil && __typeRefForAlbum.Equals(other.TypeRef()) && s.Ref() == other.Ref()
}

func (s Album) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s Album) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeRefForAlbum.Chunks()...)
	chunks = append(chunks, s._Photos.Chunks()...)
	return
}

func (s Album) Id() string {
	return s._Id
}

func (s Album) SetId(val string) Album {
	s._Id = val
	s.ref = &ref.Ref{}
	return s
}

func (s Album) Title() string {
	return s._Title
}

func (s Album) SetTitle(val string) Album {
	s._Title = val
	s.ref = &ref.Ref{}
	return s
}

func (s Album) NumPhotos() uint32 {
	return s._NumPhotos
}

func (s Album) SetNumPhotos(val uint32) Album {
	s._NumPhotos = val
	s.ref = &ref.Ref{}
	return s
}

func (s Album) Photos() RefOfSetOfRefOfRemotePhoto {
	return s._Photos
}

func (s Album) SetPhotos(val RefOfSetOfRefOfRemotePhoto) Album {
	s._Photos = val
	s.ref = &ref.Ref{}
	return s
}

// MapOfStringToAlbum

type MapOfStringToAlbum struct {
	m   types.Map
	ref *ref.Ref
}

func NewMapOfStringToAlbum() MapOfStringToAlbum {
	return MapOfStringToAlbum{types.NewMap(), &ref.Ref{}}
}

type MapOfStringToAlbumDef map[string]AlbumDef

func (def MapOfStringToAlbumDef) New() MapOfStringToAlbum {
	kv := make([]types.Value, 0, len(def)*2)
	for k, v := range def {
		kv = append(kv, types.NewString(k), v.New())
	}
	return MapOfStringToAlbum{types.NewMap(kv...), &ref.Ref{}}
}

func (m MapOfStringToAlbum) Def() MapOfStringToAlbumDef {
	def := make(map[string]AlbumDef)
	m.m.Iter(func(k, v types.Value) bool {
		def[k.(types.String).String()] = v.(Album).Def()
		return false
	})
	return def
}

func (m MapOfStringToAlbum) Equals(other types.Value) bool {
	return other != nil && __typeRefForMapOfStringToAlbum.Equals(other.TypeRef()) && m.Ref() == other.Ref()
}

func (m MapOfStringToAlbum) Ref() ref.Ref {
	return types.EnsureRef(m.ref, m)
}

func (m MapOfStringToAlbum) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, m.TypeRef().Chunks()...)
	chunks = append(chunks, m.m.Chunks()...)
	return
}

// A Noms Value that describes MapOfStringToAlbum.
var __typeRefForMapOfStringToAlbum types.TypeRef

func (m MapOfStringToAlbum) TypeRef() types.TypeRef {
	return __typeRefForMapOfStringToAlbum
}

func init() {
	__typeRefForMapOfStringToAlbum = types.MakeCompoundTypeRef(types.MapKind, types.MakePrimitiveTypeRef(types.StringKind), types.MakeTypeRef(__mainPackageInFile_picasa_CachedRef, 1))
	types.RegisterValue(__typeRefForMapOfStringToAlbum, builderForMapOfStringToAlbum, readerForMapOfStringToAlbum)
}

func builderForMapOfStringToAlbum(v types.Value) types.Value {
	return MapOfStringToAlbum{v.(types.Map), &ref.Ref{}}
}

func readerForMapOfStringToAlbum(v types.Value) types.Value {
	return v.(MapOfStringToAlbum).m
}

func (m MapOfStringToAlbum) Empty() bool {
	return m.m.Empty()
}

func (m MapOfStringToAlbum) Len() uint64 {
	return m.m.Len()
}

func (m MapOfStringToAlbum) Has(p string) bool {
	return m.m.Has(types.NewString(p))
}

func (m MapOfStringToAlbum) Get(p string) Album {
	return m.m.Get(types.NewString(p)).(Album)
}

func (m MapOfStringToAlbum) MaybeGet(p string) (Album, bool) {
	v, ok := m.m.MaybeGet(types.NewString(p))
	if !ok {
		return NewAlbum(), false
	}
	return v.(Album), ok
}

func (m MapOfStringToAlbum) Set(k string, v Album) MapOfStringToAlbum {
	return MapOfStringToAlbum{m.m.Set(types.NewString(k), v), &ref.Ref{}}
}

// TODO: Implement SetM?

func (m MapOfStringToAlbum) Remove(p string) MapOfStringToAlbum {
	return MapOfStringToAlbum{m.m.Remove(types.NewString(p)), &ref.Ref{}}
}

type MapOfStringToAlbumIterCallback func(k string, v Album) (stop bool)

func (m MapOfStringToAlbum) Iter(cb MapOfStringToAlbumIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), v.(Album))
	})
}

type MapOfStringToAlbumIterAllCallback func(k string, v Album)

func (m MapOfStringToAlbum) IterAll(cb MapOfStringToAlbumIterAllCallback) {
	m.m.IterAll(func(k, v types.Value) {
		cb(k.(types.String).String(), v.(Album))
	})
}

type MapOfStringToAlbumFilterCallback func(k string, v Album) (keep bool)

func (m MapOfStringToAlbum) Filter(cb MapOfStringToAlbumFilterCallback) MapOfStringToAlbum {
	out := m.m.Filter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), v.(Album))
	})
	return MapOfStringToAlbum{out, &ref.Ref{}}
}

// SetOfRemotePhoto

type SetOfRemotePhoto struct {
	s   types.Set
	ref *ref.Ref
}

func NewSetOfRemotePhoto() SetOfRemotePhoto {
	return SetOfRemotePhoto{types.NewSet(), &ref.Ref{}}
}

func (s SetOfRemotePhoto) Equals(other types.Value) bool {
	return other != nil && __typeRefForSetOfRemotePhoto.Equals(other.TypeRef()) && s.Ref() == other.Ref()
}

func (s SetOfRemotePhoto) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s SetOfRemotePhoto) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, s.TypeRef().Chunks()...)
	chunks = append(chunks, s.s.Chunks()...)
	return
}

// A Noms Value that describes SetOfRemotePhoto.
var __typeRefForSetOfRemotePhoto types.TypeRef

func (m SetOfRemotePhoto) TypeRef() types.TypeRef {
	return __typeRefForSetOfRemotePhoto
}

func init() {
	__typeRefForSetOfRemotePhoto = types.MakeCompoundTypeRef(types.SetKind, types.MakeTypeRef(ref.Parse("sha1-00419ebbb418539af67238164b20341913efeb4d"), 0))
	types.RegisterValue(__typeRefForSetOfRemotePhoto, builderForSetOfRemotePhoto, readerForSetOfRemotePhoto)
}

func builderForSetOfRemotePhoto(v types.Value) types.Value {
	return SetOfRemotePhoto{v.(types.Set), &ref.Ref{}}
}

func readerForSetOfRemotePhoto(v types.Value) types.Value {
	return v.(SetOfRemotePhoto).s
}

func (s SetOfRemotePhoto) Empty() bool {
	return s.s.Empty()
}

func (s SetOfRemotePhoto) Len() uint64 {
	return s.s.Len()
}

func (s SetOfRemotePhoto) Has(p RemotePhoto) bool {
	return s.s.Has(p)
}

type SetOfRemotePhotoIterCallback func(p RemotePhoto) (stop bool)

func (s SetOfRemotePhoto) Iter(cb SetOfRemotePhotoIterCallback) {
	s.s.Iter(func(v types.Value) bool {
		return cb(v.(RemotePhoto))
	})
}

type SetOfRemotePhotoIterAllCallback func(p RemotePhoto)

func (s SetOfRemotePhoto) IterAll(cb SetOfRemotePhotoIterAllCallback) {
	s.s.IterAll(func(v types.Value) {
		cb(v.(RemotePhoto))
	})
}

func (s SetOfRemotePhoto) IterAllP(concurrency int, cb SetOfRemotePhotoIterAllCallback) {
	s.s.IterAllP(concurrency, func(v types.Value) {
		cb(v.(RemotePhoto))
	})
}

type SetOfRemotePhotoFilterCallback func(p RemotePhoto) (keep bool)

func (s SetOfRemotePhoto) Filter(cb SetOfRemotePhotoFilterCallback) SetOfRemotePhoto {
	out := s.s.Filter(func(v types.Value) bool {
		return cb(v.(RemotePhoto))
	})
	return SetOfRemotePhoto{out, &ref.Ref{}}
}

func (s SetOfRemotePhoto) Insert(p ...RemotePhoto) SetOfRemotePhoto {
	return SetOfRemotePhoto{s.s.Insert(s.fromElemSlice(p)...), &ref.Ref{}}
}

func (s SetOfRemotePhoto) Remove(p ...RemotePhoto) SetOfRemotePhoto {
	return SetOfRemotePhoto{s.s.Remove(s.fromElemSlice(p)...), &ref.Ref{}}
}

func (s SetOfRemotePhoto) Union(others ...SetOfRemotePhoto) SetOfRemotePhoto {
	return SetOfRemotePhoto{s.s.Union(s.fromStructSlice(others)...), &ref.Ref{}}
}

func (s SetOfRemotePhoto) Subtract(others ...SetOfRemotePhoto) SetOfRemotePhoto {
	return SetOfRemotePhoto{s.s.Subtract(s.fromStructSlice(others)...), &ref.Ref{}}
}

func (s SetOfRemotePhoto) Any() RemotePhoto {
	return s.s.Any().(RemotePhoto)
}

func (s SetOfRemotePhoto) fromStructSlice(p []SetOfRemotePhoto) []types.Set {
	r := make([]types.Set, len(p))
	for i, v := range p {
		r[i] = v.s
	}
	return r
}

func (s SetOfRemotePhoto) fromElemSlice(p []RemotePhoto) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v
	}
	return r
}

// RefOfUser

type RefOfUser struct {
	target ref.Ref
	ref    *ref.Ref
}

func NewRefOfUser(target ref.Ref) RefOfUser {
	return RefOfUser{target, &ref.Ref{}}
}

func (r RefOfUser) TargetRef() ref.Ref {
	return r.target
}

func (r RefOfUser) Ref() ref.Ref {
	return types.EnsureRef(r.ref, r)
}

func (r RefOfUser) Equals(other types.Value) bool {
	return other != nil && __typeRefForRefOfUser.Equals(other.TypeRef()) && r.Ref() == other.Ref()
}

func (r RefOfUser) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, r.TypeRef().Chunks()...)
	chunks = append(chunks, r.target)
	return
}

// A Noms Value that describes RefOfUser.
var __typeRefForRefOfUser types.TypeRef

func (m RefOfUser) TypeRef() types.TypeRef {
	return __typeRefForRefOfUser
}

func init() {
	__typeRefForRefOfUser = types.MakeCompoundTypeRef(types.RefKind, types.MakeTypeRef(__mainPackageInFile_picasa_CachedRef, 0))
	types.RegisterRef(__typeRefForRefOfUser, builderForRefOfUser)
}

func builderForRefOfUser(r ref.Ref) types.Value {
	return NewRefOfUser(r)
}

func (r RefOfUser) TargetValue(cs chunks.ChunkSource) User {
	return types.ReadValue(r.target, cs).(User)
}

func (r RefOfUser) SetTargetValue(val User, cs chunks.ChunkSink) RefOfUser {
	return NewRefOfUser(types.WriteValue(val, cs))
}

// RefOfSetOfRefOfRemotePhoto

type RefOfSetOfRefOfRemotePhoto struct {
	target ref.Ref
	ref    *ref.Ref
}

func NewRefOfSetOfRefOfRemotePhoto(target ref.Ref) RefOfSetOfRefOfRemotePhoto {
	return RefOfSetOfRefOfRemotePhoto{target, &ref.Ref{}}
}

func (r RefOfSetOfRefOfRemotePhoto) TargetRef() ref.Ref {
	return r.target
}

func (r RefOfSetOfRefOfRemotePhoto) Ref() ref.Ref {
	return types.EnsureRef(r.ref, r)
}

func (r RefOfSetOfRefOfRemotePhoto) Equals(other types.Value) bool {
	return other != nil && __typeRefForRefOfSetOfRefOfRemotePhoto.Equals(other.TypeRef()) && r.Ref() == other.Ref()
}

func (r RefOfSetOfRefOfRemotePhoto) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, r.TypeRef().Chunks()...)
	chunks = append(chunks, r.target)
	return
}

// A Noms Value that describes RefOfSetOfRefOfRemotePhoto.
var __typeRefForRefOfSetOfRefOfRemotePhoto types.TypeRef

func (m RefOfSetOfRefOfRemotePhoto) TypeRef() types.TypeRef {
	return __typeRefForRefOfSetOfRefOfRemotePhoto
}

func init() {
	__typeRefForRefOfSetOfRefOfRemotePhoto = types.MakeCompoundTypeRef(types.RefKind, types.MakeCompoundTypeRef(types.SetKind, types.MakeCompoundTypeRef(types.RefKind, types.MakeTypeRef(ref.Parse("sha1-00419ebbb418539af67238164b20341913efeb4d"), 0))))
	types.RegisterRef(__typeRefForRefOfSetOfRefOfRemotePhoto, builderForRefOfSetOfRefOfRemotePhoto)
}

func builderForRefOfSetOfRefOfRemotePhoto(r ref.Ref) types.Value {
	return NewRefOfSetOfRefOfRemotePhoto(r)
}

func (r RefOfSetOfRefOfRemotePhoto) TargetValue(cs chunks.ChunkSource) SetOfRefOfRemotePhoto {
	return types.ReadValue(r.target, cs).(SetOfRefOfRemotePhoto)
}

func (r RefOfSetOfRefOfRemotePhoto) SetTargetValue(val SetOfRefOfRemotePhoto, cs chunks.ChunkSink) RefOfSetOfRefOfRemotePhoto {
	return NewRefOfSetOfRefOfRemotePhoto(types.WriteValue(val, cs))
}

// SetOfRefOfRemotePhoto

type SetOfRefOfRemotePhoto struct {
	s   types.Set
	ref *ref.Ref
}

func NewSetOfRefOfRemotePhoto() SetOfRefOfRemotePhoto {
	return SetOfRefOfRemotePhoto{types.NewSet(), &ref.Ref{}}
}

type SetOfRefOfRemotePhotoDef map[ref.Ref]bool

func (def SetOfRefOfRemotePhotoDef) New() SetOfRefOfRemotePhoto {
	l := make([]types.Value, len(def))
	i := 0
	for d, _ := range def {
		l[i] = NewRefOfRemotePhoto(d)
		i++
	}
	return SetOfRefOfRemotePhoto{types.NewSet(l...), &ref.Ref{}}
}

func (s SetOfRefOfRemotePhoto) Def() SetOfRefOfRemotePhotoDef {
	def := make(map[ref.Ref]bool, s.Len())
	s.s.Iter(func(v types.Value) bool {
		def[v.(RefOfRemotePhoto).TargetRef()] = true
		return false
	})
	return def
}

func (s SetOfRefOfRemotePhoto) Equals(other types.Value) bool {
	return other != nil && __typeRefForSetOfRefOfRemotePhoto.Equals(other.TypeRef()) && s.Ref() == other.Ref()
}

func (s SetOfRefOfRemotePhoto) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s SetOfRefOfRemotePhoto) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, s.TypeRef().Chunks()...)
	chunks = append(chunks, s.s.Chunks()...)
	return
}

// A Noms Value that describes SetOfRefOfRemotePhoto.
var __typeRefForSetOfRefOfRemotePhoto types.TypeRef

func (m SetOfRefOfRemotePhoto) TypeRef() types.TypeRef {
	return __typeRefForSetOfRefOfRemotePhoto
}

func init() {
	__typeRefForSetOfRefOfRemotePhoto = types.MakeCompoundTypeRef(types.SetKind, types.MakeCompoundTypeRef(types.RefKind, types.MakeTypeRef(ref.Parse("sha1-00419ebbb418539af67238164b20341913efeb4d"), 0)))
	types.RegisterValue(__typeRefForSetOfRefOfRemotePhoto, builderForSetOfRefOfRemotePhoto, readerForSetOfRefOfRemotePhoto)
}

func builderForSetOfRefOfRemotePhoto(v types.Value) types.Value {
	return SetOfRefOfRemotePhoto{v.(types.Set), &ref.Ref{}}
}

func readerForSetOfRefOfRemotePhoto(v types.Value) types.Value {
	return v.(SetOfRefOfRemotePhoto).s
}

func (s SetOfRefOfRemotePhoto) Empty() bool {
	return s.s.Empty()
}

func (s SetOfRefOfRemotePhoto) Len() uint64 {
	return s.s.Len()
}

func (s SetOfRefOfRemotePhoto) Has(p RefOfRemotePhoto) bool {
	return s.s.Has(p)
}

type SetOfRefOfRemotePhotoIterCallback func(p RefOfRemotePhoto) (stop bool)

func (s SetOfRefOfRemotePhoto) Iter(cb SetOfRefOfRemotePhotoIterCallback) {
	s.s.Iter(func(v types.Value) bool {
		return cb(v.(RefOfRemotePhoto))
	})
}

type SetOfRefOfRemotePhotoIterAllCallback func(p RefOfRemotePhoto)

func (s SetOfRefOfRemotePhoto) IterAll(cb SetOfRefOfRemotePhotoIterAllCallback) {
	s.s.IterAll(func(v types.Value) {
		cb(v.(RefOfRemotePhoto))
	})
}

func (s SetOfRefOfRemotePhoto) IterAllP(concurrency int, cb SetOfRefOfRemotePhotoIterAllCallback) {
	s.s.IterAllP(concurrency, func(v types.Value) {
		cb(v.(RefOfRemotePhoto))
	})
}

type SetOfRefOfRemotePhotoFilterCallback func(p RefOfRemotePhoto) (keep bool)

func (s SetOfRefOfRemotePhoto) Filter(cb SetOfRefOfRemotePhotoFilterCallback) SetOfRefOfRemotePhoto {
	out := s.s.Filter(func(v types.Value) bool {
		return cb(v.(RefOfRemotePhoto))
	})
	return SetOfRefOfRemotePhoto{out, &ref.Ref{}}
}

func (s SetOfRefOfRemotePhoto) Insert(p ...RefOfRemotePhoto) SetOfRefOfRemotePhoto {
	return SetOfRefOfRemotePhoto{s.s.Insert(s.fromElemSlice(p)...), &ref.Ref{}}
}

func (s SetOfRefOfRemotePhoto) Remove(p ...RefOfRemotePhoto) SetOfRefOfRemotePhoto {
	return SetOfRefOfRemotePhoto{s.s.Remove(s.fromElemSlice(p)...), &ref.Ref{}}
}

func (s SetOfRefOfRemotePhoto) Union(others ...SetOfRefOfRemotePhoto) SetOfRefOfRemotePhoto {
	return SetOfRefOfRemotePhoto{s.s.Union(s.fromStructSlice(others)...), &ref.Ref{}}
}

func (s SetOfRefOfRemotePhoto) Subtract(others ...SetOfRefOfRemotePhoto) SetOfRefOfRemotePhoto {
	return SetOfRefOfRemotePhoto{s.s.Subtract(s.fromStructSlice(others)...), &ref.Ref{}}
}

func (s SetOfRefOfRemotePhoto) Any() RefOfRemotePhoto {
	return s.s.Any().(RefOfRemotePhoto)
}

func (s SetOfRefOfRemotePhoto) fromStructSlice(p []SetOfRefOfRemotePhoto) []types.Set {
	r := make([]types.Set, len(p))
	for i, v := range p {
		r[i] = v.s
	}
	return r
}

func (s SetOfRefOfRemotePhoto) fromElemSlice(p []RefOfRemotePhoto) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v
	}
	return r
}

// RefOfRemotePhoto

type RefOfRemotePhoto struct {
	target ref.Ref
	ref    *ref.Ref
}

func NewRefOfRemotePhoto(target ref.Ref) RefOfRemotePhoto {
	return RefOfRemotePhoto{target, &ref.Ref{}}
}

func (r RefOfRemotePhoto) TargetRef() ref.Ref {
	return r.target
}

func (r RefOfRemotePhoto) Ref() ref.Ref {
	return types.EnsureRef(r.ref, r)
}

func (r RefOfRemotePhoto) Equals(other types.Value) bool {
	return other != nil && __typeRefForRefOfRemotePhoto.Equals(other.TypeRef()) && r.Ref() == other.Ref()
}

func (r RefOfRemotePhoto) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, r.TypeRef().Chunks()...)
	chunks = append(chunks, r.target)
	return
}

// A Noms Value that describes RefOfRemotePhoto.
var __typeRefForRefOfRemotePhoto types.TypeRef

func (m RefOfRemotePhoto) TypeRef() types.TypeRef {
	return __typeRefForRefOfRemotePhoto
}

func init() {
	__typeRefForRefOfRemotePhoto = types.MakeCompoundTypeRef(types.RefKind, types.MakeTypeRef(ref.Parse("sha1-00419ebbb418539af67238164b20341913efeb4d"), 0))
	types.RegisterRef(__typeRefForRefOfRemotePhoto, builderForRefOfRemotePhoto)
}

func builderForRefOfRemotePhoto(r ref.Ref) types.Value {
	return NewRefOfRemotePhoto(r)
}

func (r RefOfRemotePhoto) TargetValue(cs chunks.ChunkSource) RemotePhoto {
	return types.ReadValue(r.target, cs).(RemotePhoto)
}

func (r RefOfRemotePhoto) SetTargetValue(val RemotePhoto, cs chunks.ChunkSink) RefOfRemotePhoto {
	return NewRefOfRemotePhoto(types.WriteValue(val, cs))
}