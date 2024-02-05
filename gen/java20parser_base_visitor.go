package parser // Java20Parser

import "github.com/antlr4-go/antlr/v4"

type BaseJava20ParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseJava20ParserVisitor) VisitStart_(ctx *Start_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitTypeIdentifier(ctx *TypeIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitUnqualifiedMethodIdentifier(ctx *UnqualifiedMethodIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitNumericType(ctx *NumericTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitIntegralType(ctx *IntegralTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitFloatingPointType(ctx *FloatingPointTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitReferenceType(ctx *ReferenceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitCoit(ctx *CoitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitClassType(ctx *ClassTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitInterfaceType(ctx *InterfaceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitTypeVariable(ctx *TypeVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitArrayType(ctx *ArrayTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitDims(ctx *DimsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitTypeParameter(ctx *TypeParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitTypeParameterModifier(ctx *TypeParameterModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitTypeBound(ctx *TypeBoundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitAdditionalBound(ctx *AdditionalBoundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitTypeArguments(ctx *TypeArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitTypeArgumentList(ctx *TypeArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitTypeArgument(ctx *TypeArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitWildcard(ctx *WildcardContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitWildcardBounds(ctx *WildcardBoundsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitModuleName(ctx *ModuleNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitPackageName(ctx *PackageNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitTypeName(ctx *TypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitPackageOrTypeName(ctx *PackageOrTypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitExpressionName(ctx *ExpressionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitMethodName(ctx *MethodNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitAmbiguousName(ctx *AmbiguousNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitCompilationUnit(ctx *CompilationUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitOrdinaryCompilationUnit(ctx *OrdinaryCompilationUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitModularCompilationUnit(ctx *ModularCompilationUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitPackageDeclaration(ctx *PackageDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitPackageModifier(ctx *PackageModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitImportDeclaration(ctx *ImportDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitSingleTypeImportDeclaration(ctx *SingleTypeImportDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitTypeImportOnDemandDeclaration(ctx *TypeImportOnDemandDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitSingleStaticImportDeclaration(ctx *SingleStaticImportDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitStaticImportOnDemandDeclaration(ctx *StaticImportOnDemandDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitTopLevelClassOrInterfaceDeclaration(ctx *TopLevelClassOrInterfaceDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitModuleDeclaration(ctx *ModuleDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitModuleDirective(ctx *ModuleDirectiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitRequiresModifier(ctx *RequiresModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitClassDeclaration(ctx *ClassDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitNormalClassDeclaration(ctx *NormalClassDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitClassModifier(ctx *ClassModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitTypeParameters(ctx *TypeParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitTypeParameterList(ctx *TypeParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitClassExtends(ctx *ClassExtendsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitClassImplements(ctx *ClassImplementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitInterfaceTypeList(ctx *InterfaceTypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitClassPermits(ctx *ClassPermitsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitClassBody(ctx *ClassBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitClassBodyDeclaration(ctx *ClassBodyDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitClassMemberDeclaration(ctx *ClassMemberDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitFieldDeclaration(ctx *FieldDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitFieldModifier(ctx *FieldModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitVariableDeclaratorList(ctx *VariableDeclaratorListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitVariableDeclarator(ctx *VariableDeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitVariableDeclaratorId(ctx *VariableDeclaratorIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitVariableInitializer(ctx *VariableInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitUnannType(ctx *UnannTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitUnannPrimitiveType(ctx *UnannPrimitiveTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitUnannReferenceType(ctx *UnannReferenceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitUnannClassOrInterfaceType(ctx *UnannClassOrInterfaceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitUCOIT(ctx *UCOITContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitUnannClassType(ctx *UnannClassTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitUnannInterfaceType(ctx *UnannInterfaceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitUnannTypeVariable(ctx *UnannTypeVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitUnannArrayType(ctx *UnannArrayTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitMethodDeclaration(ctx *MethodDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitMethodModifier(ctx *MethodModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitMethodHeader(ctx *MethodHeaderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitResult(ctx *ResultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitMethodDeclarator(ctx *MethodDeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitReceiverParameter(ctx *ReceiverParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitFormalParameterList(ctx *FormalParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitFormalParameter(ctx *FormalParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitVariableArityParameter(ctx *VariableArityParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitVariableModifier(ctx *VariableModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitThrowsT(ctx *ThrowsTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitExceptionTypeList(ctx *ExceptionTypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitExceptionType(ctx *ExceptionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitMethodBody(ctx *MethodBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitInstanceInitializer(ctx *InstanceInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitStaticInitializer(ctx *StaticInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitConstructorDeclaration(ctx *ConstructorDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitConstructorModifier(ctx *ConstructorModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitConstructorDeclarator(ctx *ConstructorDeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitSimpleTypeName(ctx *SimpleTypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitConstructorBody(ctx *ConstructorBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitExplicitConstructorInvocation(ctx *ExplicitConstructorInvocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitEnumDeclaration(ctx *EnumDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitEnumBody(ctx *EnumBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitEnumConstantList(ctx *EnumConstantListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitEnumConstant(ctx *EnumConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitEnumConstantModifier(ctx *EnumConstantModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitEnumBodyDeclarations(ctx *EnumBodyDeclarationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitRecordDeclaration(ctx *RecordDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitRecordHeader(ctx *RecordHeaderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitRecordComponentList(ctx *RecordComponentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitRecordComponent(ctx *RecordComponentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitVariableArityRecordComponent(ctx *VariableArityRecordComponentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitRecordComponentModifier(ctx *RecordComponentModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitRecordBody(ctx *RecordBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitRecordBodyDeclaration(ctx *RecordBodyDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitCompactConstructorDeclaration(ctx *CompactConstructorDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitInterfaceDeclaration(ctx *InterfaceDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitNormalInterfaceDeclaration(ctx *NormalInterfaceDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitInterfaceModifier(ctx *InterfaceModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitInterfaceExtends(ctx *InterfaceExtendsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitInterfacePermits(ctx *InterfacePermitsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitInterfaceBody(ctx *InterfaceBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitInterfaceMemberDeclaration(ctx *InterfaceMemberDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitConstantDeclaration(ctx *ConstantDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitConstantModifier(ctx *ConstantModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitInterfaceMethodDeclaration(ctx *InterfaceMethodDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitInterfaceMethodModifier(ctx *InterfaceMethodModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitAnnotationInterfaceDeclaration(ctx *AnnotationInterfaceDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitAnnotationInterfaceBody(ctx *AnnotationInterfaceBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitAnnotationInterfaceMemberDeclaration(ctx *AnnotationInterfaceMemberDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitAnnotationInterfaceElementDeclaration(ctx *AnnotationInterfaceElementDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitAnnotationInterfaceElementModifier(ctx *AnnotationInterfaceElementModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitDefaultValue(ctx *DefaultValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitAnnotation(ctx *AnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitNormalAnnotation(ctx *NormalAnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitElementValuePairList(ctx *ElementValuePairListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitElementValuePair(ctx *ElementValuePairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitElementValue(ctx *ElementValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitElementValueArrayInitializer(ctx *ElementValueArrayInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitElementValueList(ctx *ElementValueListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitMarkerAnnotation(ctx *MarkerAnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitSingleElementAnnotation(ctx *SingleElementAnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitArrayInitializer(ctx *ArrayInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitVariableInitializerList(ctx *VariableInitializerListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitBlockStatements(ctx *BlockStatementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitBlockStatement(ctx *BlockStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitLocalClassOrInterfaceDeclaration(ctx *LocalClassOrInterfaceDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitLocalVariableType(ctx *LocalVariableTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitLocalVariableDeclarationStatement(ctx *LocalVariableDeclarationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitStatementNoShortIf(ctx *StatementNoShortIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitStatementWithoutTrailingSubstatement(ctx *StatementWithoutTrailingSubstatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitEmptyStatement_(ctx *EmptyStatement_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitLabeledStatement(ctx *LabeledStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitLabeledStatementNoShortIf(ctx *LabeledStatementNoShortIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitExpressionStatement(ctx *ExpressionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitStatementExpression(ctx *StatementExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitIfThenStatement(ctx *IfThenStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitIfThenElseStatement(ctx *IfThenElseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitIfThenElseStatementNoShortIf(ctx *IfThenElseStatementNoShortIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitAssertStatement(ctx *AssertStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitSwitchStatement(ctx *SwitchStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitSwitchBlock(ctx *SwitchBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitSwitchRule(ctx *SwitchRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitSwitchBlockStatementGroup(ctx *SwitchBlockStatementGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitSwitchLabel(ctx *SwitchLabelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitCaseConstant(ctx *CaseConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitWhileStatementNoShortIf(ctx *WhileStatementNoShortIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitDoStatement(ctx *DoStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitForStatement(ctx *ForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitForStatementNoShortIf(ctx *ForStatementNoShortIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitBasicForStatement(ctx *BasicForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitBasicForStatementNoShortIf(ctx *BasicForStatementNoShortIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitForInit(ctx *ForInitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitForUpdate(ctx *ForUpdateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitStatementExpressionList(ctx *StatementExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitEnhancedForStatement(ctx *EnhancedForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitEnhancedForStatementNoShortIf(ctx *EnhancedForStatementNoShortIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitContinueStatement(ctx *ContinueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitThrowStatement(ctx *ThrowStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitSynchronizedStatement(ctx *SynchronizedStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitTryStatement1(ctx *TryStatement1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitTryStatement2(ctx *TryStatement2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitTryStatement3(ctx *TryStatement3Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitTryStatement4(ctx *TryStatement4Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitCatches(ctx *CatchesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitCatchClause(ctx *CatchClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitCatchFormalParameter(ctx *CatchFormalParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitCatchType(ctx *CatchTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitFinallyBlock(ctx *FinallyBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitTryWithResourcesStatement(ctx *TryWithResourcesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitResourceSpecification(ctx *ResourceSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitResourceList(ctx *ResourceListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitResource(ctx *ResourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitVariableAccess(ctx *VariableAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitYieldStatement(ctx *YieldStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitPattern(ctx *PatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitTypePattern(ctx *TypePatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitPrimary(ctx *PrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitPrimaryNoNewArray(ctx *PrimaryNoNewArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitPNNA(ctx *PNNAContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitClassLiteral(ctx *ClassLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitClassInstanceCreationExpression(ctx *ClassInstanceCreationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitUnqualifiedClassInstanceCreationExpression(ctx *UnqualifiedClassInstanceCreationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitClassOrInterfaceTypeToInstantiate(ctx *ClassOrInterfaceTypeToInstantiateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitTypeArgumentsOrDiamond(ctx *TypeArgumentsOrDiamondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitArrayCreationExpression(ctx *ArrayCreationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitArrayCreationExpressionWithoutInitializer(ctx *ArrayCreationExpressionWithoutInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitArrayCreationExpressionWithInitializer(ctx *ArrayCreationExpressionWithInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitDimExprs(ctx *DimExprsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitDimExpr(ctx *DimExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitArrayAccess(ctx *ArrayAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitFieldAccess(ctx *FieldAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitMethodInvocation(ctx *MethodInvocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitArgumentList(ctx *ArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitMethodReference(ctx *MethodReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitPostfixExpression(ctx *PostfixExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitPfE(ctx *PfEContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitPostIncrementExpression(ctx *PostIncrementExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitPostDecrementExpression(ctx *PostDecrementExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitUnaryExpression(ctx *UnaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitPreIncrementExpression(ctx *PreIncrementExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitPreDecrementExpression(ctx *PreDecrementExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitUnaryExpressionNotPlusMinus(ctx *UnaryExpressionNotPlusMinusContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitCastExpression(ctx *CastExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitShiftExpression(ctx *ShiftExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitRelationalExpression(ctx *RelationalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitEqualityExpression(ctx *EqualityExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitAndExpression(ctx *AndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitExclusiveOrExpression(ctx *ExclusiveOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitInclusiveOrExpression(ctx *InclusiveOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitConditionalAndExpression(ctx *ConditionalAndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitConditionalOrExpression(ctx *ConditionalOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitConditionalExpression(ctx *ConditionalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitLeftHandSide(ctx *LeftHandSideContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitLambdaExpression(ctx *LambdaExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitLambdaParameters(ctx *LambdaParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitLambdaParameterList(ctx *LambdaParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitLambdaParameter(ctx *LambdaParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitLambdaParameterType(ctx *LambdaParameterTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitLambdaBody(ctx *LambdaBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitSwitchExpression(ctx *SwitchExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava20ParserVisitor) VisitConstantExpression(ctx *ConstantExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}
