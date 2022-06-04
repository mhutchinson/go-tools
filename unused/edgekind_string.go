// Code generated by "stringer -type EdgeKind"; DO NOT EDIT.

package unused

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[edgeAlias-1]
	_ = x[edgeBlankField-2]
	_ = x[edgeAnonymousStruct-4]
	_ = x[edgeCgoExported-8]
	_ = x[edgeConstGroup-16]
	_ = x[edgeElementType-32]
	_ = x[edgeEmbeddedInterface-64]
	_ = x[edgeExportedConstant-128]
	_ = x[edgeExportedField-256]
	_ = x[edgeExportedFunction-512]
	_ = x[edgeExportedMethod-1024]
	_ = x[edgeExportedType-2048]
	_ = x[edgeExportedVariable-4096]
	_ = x[edgeExtendsExportedFields-8192]
	_ = x[edgeExtendsExportedMethodSet-16384]
	_ = x[edgeFieldAccess-32768]
	_ = x[edgeFunctionArgument-65536]
	_ = x[edgeFunctionResult-131072]
	_ = x[edgeFunctionSignature-262144]
	_ = x[edgeImplements-524288]
	_ = x[edgeInstructionOperand-1048576]
	_ = x[edgeInterfaceCall-2097152]
	_ = x[edgeInterfaceMethod-4194304]
	_ = x[edgeKeyType-8388608]
	_ = x[edgeLinkname-16777216]
	_ = x[edgeMainFunction-33554432]
	_ = x[edgeNamedType-67108864]
	_ = x[edgeNoCopySentinel-134217728]
	_ = x[edgeProvidesMethod-268435456]
	_ = x[edgeReceiver-536870912]
	_ = x[edgeRuntimeFunction-1073741824]
	_ = x[edgeSignature-2147483648]
	_ = x[edgeStructConversion-4294967296]
	_ = x[edgeTestSink-8589934592]
	_ = x[edgeTupleElement-17179869184]
	_ = x[edgeType-34359738368]
	_ = x[edgeTypeName-68719476736]
	_ = x[edgeUnderlyingType-137438953472]
	_ = x[edgePointerType-274877906944]
	_ = x[edgeUnsafeConversion-549755813888]
	_ = x[edgeUsedConstant-1099511627776]
	_ = x[edgeVarDecl-2199023255552]
	_ = x[edgeIgnored-4398046511104]
	_ = x[edgeSamePointer-8796093022208]
	_ = x[edgeTypeParam-17592186044416]
	_ = x[edgeTypeArg-35184372088832]
	_ = x[edgeUnionTerm-70368744177664]
}

const _EdgeKind_name = "edgeAliasedgeBlankFieldedgeAnonymousStructedgeCgoExportededgeConstGroupedgeElementTypeedgeEmbeddedInterfaceedgeExportedConstantedgeExportedFieldedgeExportedFunctionedgeExportedMethodedgeExportedTypeedgeExportedVariableedgeExtendsExportedFieldsedgeExtendsExportedMethodSetedgeFieldAccessedgeFunctionArgumentedgeFunctionResultedgeFunctionSignatureedgeImplementsedgeInstructionOperandedgeInterfaceCalledgeInterfaceMethodedgeKeyTypeedgeLinknameedgeMainFunctionedgeNamedTypeedgeNoCopySentineledgeProvidesMethodedgeReceiveredgeRuntimeFunctionedgeSignatureedgeStructConversionedgeTestSinkedgeTupleElementedgeTypeedgeTypeNameedgeUnderlyingTypeedgePointerTypeedgeUnsafeConversionedgeUsedConstantedgeVarDecledgeIgnorededgeSamePointeredgeTypeParamedgeTypeArgedgeUnionTerm"

var _EdgeKind_map = map[EdgeKind]string{
	1:              _EdgeKind_name[0:9],
	2:              _EdgeKind_name[9:23],
	4:              _EdgeKind_name[23:42],
	8:              _EdgeKind_name[42:57],
	16:             _EdgeKind_name[57:71],
	32:             _EdgeKind_name[71:86],
	64:             _EdgeKind_name[86:107],
	128:            _EdgeKind_name[107:127],
	256:            _EdgeKind_name[127:144],
	512:            _EdgeKind_name[144:164],
	1024:           _EdgeKind_name[164:182],
	2048:           _EdgeKind_name[182:198],
	4096:           _EdgeKind_name[198:218],
	8192:           _EdgeKind_name[218:243],
	16384:          _EdgeKind_name[243:271],
	32768:          _EdgeKind_name[271:286],
	65536:          _EdgeKind_name[286:306],
	131072:         _EdgeKind_name[306:324],
	262144:         _EdgeKind_name[324:345],
	524288:         _EdgeKind_name[345:359],
	1048576:        _EdgeKind_name[359:381],
	2097152:        _EdgeKind_name[381:398],
	4194304:        _EdgeKind_name[398:417],
	8388608:        _EdgeKind_name[417:428],
	16777216:       _EdgeKind_name[428:440],
	33554432:       _EdgeKind_name[440:456],
	67108864:       _EdgeKind_name[456:469],
	134217728:      _EdgeKind_name[469:487],
	268435456:      _EdgeKind_name[487:505],
	536870912:      _EdgeKind_name[505:517],
	1073741824:     _EdgeKind_name[517:536],
	2147483648:     _EdgeKind_name[536:549],
	4294967296:     _EdgeKind_name[549:569],
	8589934592:     _EdgeKind_name[569:581],
	17179869184:    _EdgeKind_name[581:597],
	34359738368:    _EdgeKind_name[597:605],
	68719476736:    _EdgeKind_name[605:617],
	137438953472:   _EdgeKind_name[617:635],
	274877906944:   _EdgeKind_name[635:650],
	549755813888:   _EdgeKind_name[650:670],
	1099511627776:  _EdgeKind_name[670:686],
	2199023255552:  _EdgeKind_name[686:697],
	4398046511104:  _EdgeKind_name[697:708],
	8796093022208:  _EdgeKind_name[708:723],
	17592186044416: _EdgeKind_name[723:736],
	35184372088832: _EdgeKind_name[736:747],
	70368744177664: _EdgeKind_name[747:760],
}

func (i EdgeKind) String() string {
	if str, ok := _EdgeKind_map[i]; ok {
		return str
	}
	return "EdgeKind(" + strconv.FormatInt(int64(i), 10) + ")"
}
